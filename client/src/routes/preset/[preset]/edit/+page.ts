import { connectBackend } from '$lib/backend';
import { Preset, type PresetData } from '$lib/preset';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const id = params.preset;
	const backend = connectBackend();

	const presetRecord = await backend.app.records.getOne('presets', id);
	const preset = Preset.fromRecord(presetRecord);

	const presetDataRecords = await backend.app.records.getFullList('preset_data', undefined, {
		filter: `preset='${id}'`
	});
	const presetData: Omit<PresetData, 'preset'>[] = presetDataRecords.map((record) => ({
		id: record.id,
		data: record.data,
		title: record.title
	}));

	const presetStats = await backend.fetchPresetStats(id);
	const authorName = (await backend.app.records.getOne('profiles', preset.author)).name;

	return { preset, presetData, presetStats, authorName };
};
