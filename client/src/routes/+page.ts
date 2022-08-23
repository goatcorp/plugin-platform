import { Backend, connectBackend } from '$lib/backend';
import { Preset, type PresetStats } from '$lib/preset';

const getNewPresets = async (backend: Backend) => {
	const presetList = await backend.app.records.getList('presets', 1, undefined, {
		sort: '-created'
	});
	const presets = presetList.items.map((item) => Preset.fromRecord(item));
	return presets;
};

const getPopularPresets = async (backend: Backend) => {
	const presetList = await backend.search('/api/presets/popular');
	const presets = presetList.items.map((item) => {
		const preset = Preset.fromRecord(item);
		return { ...preset, views: item.views };
	});
	return presets;
};

const getTrendingPresets = async (backend: Backend) => {
	const presetList = await backend.search('/api/presets/trending');
	const presets = presetList.items.map((item) => Preset.fromRecord(item));
	return presets;
};

export async function load() {
	const backend = connectBackend();

	const presets: Preset[] = (await getNewPresets(backend)).slice(0, 12);
	const popular: (Preset & PresetStats)[] = (await getPopularPresets(backend)).slice(0, 12);
	const trending: Preset[] = (await getTrendingPresets(backend)).slice(0, 12);

	const stats: Record<string, PresetStats> = {};
	const authors: Record<string, string> = {};
	for (const preset of presets.concat(popular).concat(trending)) {
		if (!authors[preset.author]) {
			const author = await backend.app.records.getOne('profiles', preset.author);
			authors[preset.author] = author.name;
		}

		if (!stats[preset.id]) {
			stats[preset.id] = await backend.fetchPresetStats(preset.id);
		}
	}

	return { presets, popular, trending, stats, authors };
}
