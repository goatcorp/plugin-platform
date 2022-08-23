import type { PresetStats } from '$lib/preset';
import type { Profile } from '$lib/profile';
import type { PageLoad } from './$types';
import { connectBackend } from '$lib/backend';

export const load: PageLoad = async ({ url, params }) => {
	const page: number = parseInt(url.searchParams.get('page') || '1');
	const id = params.user;

	const backend = connectBackend();

	const profileRecord = await backend.app.records.getOne('profiles', id);
	const profile: Profile = {
		id: profileRecord.id,
		name: profileRecord.name,
		avatar: profileRecord.avatar
	};

	const presets = await backend.fetchProfileFavoritePresetsPaginated(profile.id, page);
	const presetAuthors: Record<string, string> = {};
	const presetStats: Record<string, PresetStats> = {};

	for (const preset of presets.items) {
		if (!presetAuthors[preset.author]) {
			const author = await backend.app.records.getOne('profiles', preset.author);
			presetAuthors[preset.author] = author.name;
		}

		if (!presetStats[preset.id]) {
			presetStats[preset.id] = await backend.fetchPresetStats(preset.id);
		}
	}

	return {
		profile,
		presets: presets.items,
		presetAuthors,
		presetStats,
		page: presets.page,
		totalPages: presets.totalPages
	};
};
