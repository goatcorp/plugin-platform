import PocketBase from 'pocketbase';
import type { Preset, PresetStats } from '$lib/preset';
import type { PageLoad } from './$types';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export const load: PageLoad = async ({ url }) => {
	const page: number = parseInt(url.searchParams.get('page') || '1');

	const client = connect();
	const records = await client.records.getList('presets', page, 50);
	const presets: Preset[] = records.items.map((item) => ({
		id: item.id,
		thumbnail: item.thumbnail,
		title: item.title,
		author: item.author,
		created: new Date(item.created),
		updated: new Date(item.updated)
	}));

	const stats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		stats[preset.id] = presetStatsRecords
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

	const authors: Record<string, string> = {};
	for (const result of presets) {
		const user = await client.records.getOne('profiles', result.author);
		authors[result.author] = user.name;
	}

	return { presets, authors, stats, page: records.page, totalPages: records.totalPages };
};
