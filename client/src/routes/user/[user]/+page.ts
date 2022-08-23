import { connectBackend } from '$lib/backend';
import { Preset, type PresetStats } from '$lib/preset';
import type { Profile } from '$lib/profile';

export async function load({ params }: { params: Record<string, string> }) {
	const id = params.user;
	const backend = connectBackend();

	const profileRecord = await backend.app.records.getOne('profiles', id);
	const profile: Profile = {
		id: profileRecord.id,
		name: profileRecord.name,
		avatar: profileRecord.avatar
	};

	const presetRecords = await backend.app.records.getFullList('presets', undefined, {
		filter: `author='${id}'`
	});
	const presets: Preset[] = presetRecords
		.map((record) => Preset.fromRecord(record))
		.sort((a, b) => b.created.valueOf() - a.created.valueOf());

	const presetStats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		presetStats[preset.id] = await backend.fetchPresetStats(preset.id);
	}

	return { profile, presets, presetStats };
}
