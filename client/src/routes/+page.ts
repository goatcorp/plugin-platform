import type { Preset, PresetStats } from '$lib/preset';
import PocketBase from 'pocketbase';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

const getNewPresets = async (client: PocketBase) => {
	const presetList = await client.records.getList('presets');
	const presets = presetList.items
		.map((item) => ({
			id: item.id,
			thumbnail: item.thumbnail,
			title: item.title,
			author: item.author,
			created: new Date(item.created),
			updated: new Date(item.updated)
		}))
		.sort((a, b) => b.created.valueOf() - a.created.valueOf());
	return presets;
};

const getPopularPresets = async (client: PocketBase) => {
	const presetList: (Preset & PresetStats)[] = await client.send('/api/presets/popular', {});
	const presets = presetList
		.map((item) => ({
			id: item.id,
			thumbnail: item.thumbnail,
			title: item.title,
			author: item.author,
			views: item.views,
			created: new Date(item.created),
			updated: new Date(item.updated)
		}))
		.sort((a, b) => b.views - a.views);
	return presets;
};

export async function load() {
	const client = connect();

	const presets: Preset[] = await getNewPresets(client);
	const popular: (Preset & PresetStats)[] = await getPopularPresets(client);

	const presetStats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		presetStats[preset.id] = presetStatsRecords
			.map((record) => ({
				views: record.views
			}))
			.reduce((agg, next) => {
				agg.views += next.views;
				return agg;
			});
	}

	return { presets, popular, presetStats };
}
