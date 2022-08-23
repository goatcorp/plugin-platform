import PocketBase, { type Record as PocketBaseRecord } from 'pocketbase';
import type { Preset, PresetStats } from '$lib/preset';
import type { PageLoad } from './$types';
import type { ListResult } from '$lib/pocketbase-ext';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export const load: PageLoad = async ({ url }) => {
	const query = url.searchParams.get('q') || '';

	const client = connect();

	const records: ListResult<PocketBaseRecord> = await client.send(
		`/api/presets/search?${url.searchParams.toString()}`,
		{}
	);
	const results: Preset[] = records.items.map((item) => ({
		id: item.id,
		thumbnail: item.thumbnail,
		title: item.title,
		author: item.author,
		spoiler: item.spoiler,
		created: new Date(item.created),
		updated: new Date(item.updated)
	}));

	const stats: Record<string, PresetStats> = {};
	for (const preset of results) {
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
	for (const result of results) {
		const user = await client.records.getOne('profiles', result.author);
		authors[result.author] = user.name;
	}

	return { results, authors, stats, query, page: records.page, totalPages: records.totalPages };
};
