import PocketBase from 'pocketbase';
import type { Record as PocketBaseRecord } from 'pocketbase';
import { Preset, type PresetStats } from './preset';
import type { Tag } from './tags';
import type { Plugin } from './plugins';

interface ListResult<M> {
	page: number;
	perPage: number;
	totalItems: number;
	totalPages: number;
	items: Array<M>;
}

export class Backend {
	app: PocketBase;

	constructor(addr: string) {
		this.app = new PocketBase(addr);
	}

	/**
	 *
	 * @param path The request path, starting with `/api`.
	 * @param queryParams Any query parameters to be sent along with the request, in the URL.
	 * @param reqConfig Request configuration to be passed to `fetch` internally, with some modifications.
	 * @returns A paginated list of result records.
	 */
	search<T>(
		path: string,
		queryParams?: Record<string, string>,
		reqConfig?: Record<string, string>
	): Promise<ListResult<PocketBaseRecord & T>> {
		if (queryParams != null) {
			const p = new URLSearchParams(queryParams);
			return this.app.send(`${path}?${p.toString()}`, reqConfig || {});
		} else {
			return this.app.send(path, reqConfig || {});
		}
	}

	async addPresetTagById(presetId: string, tagId: string) {
		await this.app.records.create('preset_tags', {
			preset: presetId,
			tag: tagId
		});
	}

	async addPresetTagByLabel(presetId: string, tagLabel: string) {
		const tagsWithLabel = await this.app.records.getList('tags', 1, 1, {
			filter: `label='${tagLabel}'`
		});

		if (tagsWithLabel.items.length === 0) {
			throw new Error('Tag not found.');
		}

		await this.app.records.create('preset_tags', {
			preset: presetId,
			tag: tagsWithLabel.items[0].id
		});
	}

	async removePresetTagById(presetId: string, tagId: string) {
		const relations = await this.app.records.getFullList('preset_tags', undefined, {
			filter: `preset='${presetId}' && tag='${tagId}'`
		});

		for (const relation of relations) {
			await this.app.records.delete('preset_tags', relation.id);
		}
	}

	async removePresetTagByLabel(presetId: string, tagLabel: string) {
		const relations = await this.app.records.getFullList('preset_tags', undefined, {
			filter: `preset='${presetId}'`,
			expand: 'tag'
		});

		for (const relation of relations) {
			if (relation['@expand'].tag.label === tagLabel) {
				await this.app.records.delete('preset_tags', relation.id);
			}
		}
	}

	async searchTags(query: string): Promise<Tag[]> {
		const tags = await this.app.records.getList('tags', 1, 10, {
			filter: `label~'${query}'`
		});
		return tags.items.map((record) => ({ id: record.id, label: record.label }));
	}

	async fetchPresetStats(presetId: string): Promise<PresetStats> {
		const records = await this.app.records.getFullList('preset_stats', undefined, {
			filter: `preset='${presetId}'`
		});
		return records
			.map((record) => ({
				views: record.views
			}))
			.reduce(
				(agg, next) => {
					agg.views += next.views;
					return agg;
				},
				{
					views: 0
				}
			);
	}

	async fetchPresetTags(presetId: string): Promise<Tag[]> {
		const presetTagRelations = (
			await this.app.records.getFullList('preset_tags', undefined, {
				filter: `preset='${presetId}'`
			})
		).map((presetTag) => ({ preset: presetTag.preset, tag: presetTag.tag }));

		const presetTags: Tag[] = [];
		for (const tag of presetTagRelations) {
			const presetTag = await this.app.records.getOne('tags', tag.tag);
			presetTags.push({ id: presetTag.id, label: presetTag.label });
		}

		return presetTags;
	}

	async fetchPresetPlugin(presetId: string): Promise<Plugin | null> {
		const presetPlugins = (
			await this.app.records.getFullList('preset_plugins', undefined, {
				filter: `preset='${presetId}'`,
				expand: 'plugin'
			})
		).map((record) => {
			const { id, internal_name, name } = record['@expand'].plugin;
			return { id, internal_name, name };
		});

		return presetPlugins.length > 0 ? presetPlugins[0] : null;
	}

	async fetchProfileFavoritePresets(profileId: string): Promise<Record<string, Preset>> {
		const favorites = await this.app.records.getFullList('profile_preset_favorites', undefined, {
			filter: `profile='${profileId}'`,
			expand: 'preset'
		});

		const presets: Record<string, Preset> = {};
		for (const favorite of favorites) {
			const preset = favorite['@expand'].preset;
			presets[preset.id] = Preset.fromRecord(preset);
		}

		return presets;
	}

	async fetchProfileFavoritePresetsPaginated(
		profileId: string,
		page?: number,
		perPage?: number
	): Promise<ListResult<Preset>> {
		const favorites = await this.app.records.getList('profile_preset_favorites', page, perPage, {
			filter: `profile='${profileId}'`,
			expand: 'preset'
		});

		return {
			...favorites,
			items: favorites.items.map((item) => Preset.fromRecord(item['@expand'].preset))
		};
	}

	async isFavoritePreset(profileId: string, presetId: string): Promise<boolean> {
		const favorites = await this.app.records.getList('profile_preset_favorites', 1, 0, {
			filter: `profile='${profileId}' && preset='${presetId}'`
		});

		return favorites.totalItems > 0;
	}

	isAuthenticated() {
		return this.app.authStore.isValid;
	}

	getCurrentUser() {
		const model = this.app.authStore.model;
		if ('profile' in model) {
			return model;
		}

		return null;
	}
}

type ConnectBackend = () => Backend;

export const connectBackend: ConnectBackend = () => {
	return new Backend('http://127.0.0.1:8090');
};
