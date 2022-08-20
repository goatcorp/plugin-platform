import PocketBase from 'pocketbase';
import type { Preset, PresetData, PresetStats } from '$lib/preset';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export async function load({ params }: { params: Record<string, string> }) {
	const id = params.preset;
	const client = connect();

	const presetRecord = await client.records.getOne('presets', id);
	const preset: Preset = {
		id: presetRecord.id,
		thumbnail: presetRecord.thumbnail,
		title: presetRecord.title,
		author: presetRecord.author,
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

	const isFavoriteInitial =
		(
			await client.records.getList('user_preset_favorites', 1, 1, {
				filter: `preset='${preset.id}'`
			})
		).totalItems !== 0;

	const authorName = (await client.records.getOne('profiles', preset.author)).name;

	return { preset, presetData, presetStats, authorName, isFavoriteInitial };
}
