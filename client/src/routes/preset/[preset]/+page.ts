import PocketBase from 'pocketbase';
import type { Preset, PresetData, PresetStats } from '$lib/preset';
import type { PageLoad } from './$types';
import type { Tag } from '$lib/tags';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export const load: PageLoad = async ({ params }) => {
	const id = params.preset;
	const client = connect();

	const presetRecord = await client.records.getOne('presets', id);
	const preset: Preset = {
		id: presetRecord.id,
		thumbnail: presetRecord.thumbnail,
		title: presetRecord.title,
		author: presetRecord.author,
		spoiler: presetRecord.spoiler,
		created: new Date(presetRecord.created),
		updated: new Date(presetRecord.updated)
	};

	const presetDataRecords = await client.records.getFullList('preset_data', undefined, {
		filter: `preset~'${id}'`
	});
	const presetData: Omit<PresetData, 'preset' | 'id'>[] = presetDataRecords.map((record) => ({
		data: record.data,
		title: record.title
	}));

	const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
		filter: `preset~'${id}'`
	});
	const presetStats: PresetStats = presetStatsRecords
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

	const presetTagRelations = (
		await client.records.getFullList('preset_tags', undefined, {
			filter: `preset='${preset.id}'`
		})
	).map((presetTag) => ({ preset: presetTag.preset, tag: presetTag.tag }));
	const presetTags: Tag[] = [];
	for (const tag of presetTagRelations) {
		const presetTag = await client.records.getOne('tags', tag.tag);
		presetTags.push({ id: presetTag.id, label: presetTag.label });
	}

	const user = client.authStore.model;
	const isFavoriteInitial =
		'id' in user
			? (
					await client.records.getList('profile_preset_favorites', 1, 1, {
						filter: `profile='${preset.id}'`
					})
			  ).totalItems !== 0
			: null;

	const authorName = (await client.records.getOne('profiles', preset.author)).name;

	return { preset, presetData, presetStats, presetTags, authorName, isFavoriteInitial };
};
