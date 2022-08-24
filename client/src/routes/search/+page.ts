import { Preset, type PresetStats } from '$lib/preset';
import type { PageLoad } from './$types';
import { connectBackend } from '$lib/backend';

export const load: PageLoad = async ({ url }) => {
	const query = url.searchParams.get('q') || '';

	const backend = connectBackend();

	const searchParams: Record<string, string> = {};
	for (const [key, value] of url.searchParams.entries()) {
		searchParams[key] = value;
	}

	const records = await backend.search('/api/presets/search', searchParams);
	const results: Preset[] = records.items.map((item) => Preset.fromRecord(item));

	const stats: Record<string, PresetStats> = {};
	const authors: Record<string, string> = {};

	for (const preset of results) {
		if (!authors[preset.author]) {
			const author = await backend.app.records.getOne('profiles', preset.author);
			authors[preset.author] = author.name;
		}

		if (!stats[preset.id]) {
			stats[preset.id] = await backend.fetchPresetStats(preset.id);
		}
	}

	const pluginRecords = await backend.app.records.getFullList('plugins');
	const plugins = pluginRecords
		.map((record) => ({
			id: record.id as string,
			internal_name: record.internal_name as string,
			name: record.name as string
		}))
		.sort((a, b) => a.name.localeCompare(b.name));

	return {
		results,
		authors,
		stats,
		plugins,
		query,
		page: records.page,
		totalPages: records.totalPages
	};
};
