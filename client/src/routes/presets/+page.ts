import { Preset, type PresetStats } from '$lib/preset';
import type { PageLoad } from './$types';
import { connectBackend } from '$lib/backend';

export const load: PageLoad = async ({ url }) => {
	const backend = connectBackend();
	const records = await backend.search(
		'/api/presets/search',
		{
			page: url.searchParams.get('page') || '1',
			tags: url.searchParams.get('tags') || ''
		},
		{}
	);
	const presets: Preset[] = records.items.map((item) => Preset.fromRecord(item));

	const stats: Record<string, PresetStats> = {};
	const authors: Record<string, string> = {};
	for (const preset of presets) {
		if (!authors[preset.author]) {
			const author = await backend.app.records.getOne('profiles', preset.author);
			authors[preset.author] = author.name;
		}

		if (!stats[preset.id]) {
			stats[preset.id] = await backend.fetchPresetStats(preset.id);
		}
	}

	return { presets, authors, stats, page: records.page, totalPages: records.totalPages };
};
