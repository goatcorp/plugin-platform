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
			spoiler: item.spoiler,
			created: new Date(item.created),
			updated: new Date(item.updated)
		}))
		.sort((a, b) => b.created.valueOf() - a.created.valueOf());
	return presets;
};

const getPopularPresets = async (client: PocketBase) => {
	const presetList: (Preset & PresetStats)[] = await client.send('/api/presets/popular', {});
	const presets = presetList.map((item) => ({
		id: item.id,
		thumbnail: item.thumbnail,
		title: item.title,
		author: item.author,
		spoiler: item.spoiler,
		views: item.views,
		created: new Date(item.created),
		updated: new Date(item.updated)
	}));
	return presets;
};

const getTrendingPresets = async (client: PocketBase) => {
	const presetList: Preset[] = await client.send('/api/presets/trending', {});
	const presets = presetList.map((item) => ({
		id: item.id,
		thumbnail: item.thumbnail,
		title: item.title,
		author: item.author,
		spoiler: item.spoiler,
		created: new Date(item.created),
		updated: new Date(item.updated)
	}));
	return presets;
};

export async function load() {
	const client = connect();

	const presets: Preset[] = (await getNewPresets(client)).slice(0, 12);
	const popular: (Preset & PresetStats)[] = (await getPopularPresets(client)).slice(0, 12);
	const trending: Preset[] = (await getTrendingPresets(client)).slice(0, 12);

	const stats: Record<string, PresetStats> = {};
	const authors: Record<string, string> = {};
	for (const preset of presets.concat(popular).concat(trending)) {
		const statsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		stats[preset.id] = statsRecords
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

		const user = await client.records.getOne('profiles', preset.author);
		authors[preset.author] = user.name;
	}

	return { presets, popular, trending, stats, authors };
}
