import type { Preset, PresetStats } from '$lib/preset';
import type { Profile } from '$lib/profile';
import type { PageLoad } from './$types';
import PocketBase from 'pocketbase';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export const load: PageLoad = async ({ url, params }) => {
	const page: number = parseInt(url.searchParams.get('page') || '1');
	const id = params.user;

	const client = connect();

	const profileRecord = await client.records.getOne('profiles', id);
	const profile: Profile = {
		id: profileRecord.id,
		name: profileRecord.name,
		avatar: profileRecord.avatar
	};

	const favorites = await client.records.getList('profile_preset_favorites', page, 50, {
		filter: `profile='${id}'`
	});

	const presets: Preset[] = [];
	for (const favorite of favorites.items) {
		const preset = await client.records.getOne('presets', favorite.preset);
		presets.push({
			id: preset.id,
			thumbnail: preset.thumbnail,
			title: preset.title,
			author: preset.author,
			spoiler: preset.spoiler,
			created: new Date(preset.created),
			updated: new Date(preset.updated)
		});
	}

	const presetAuthors: Record<string, string> = {};
	for (const preset of presets) {
		const author = await client.records.getOne('profiles', preset.author);
		presetAuthors[preset.author] = author.name;
	}

	const presetStats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		presetStats[preset.id] = presetStatsRecords
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

	return {
		profile,
		presets,
		presetAuthors,
		presetStats,
		page: favorites.page,
		totalPages: favorites.totalPages
	};
};
