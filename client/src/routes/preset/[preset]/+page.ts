import PocketBase from 'pocketbase';
import type { Preset, PresetData } from '$lib/preset';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export async function load({ params }: { params: Record<string, string> }) {
	const id = params.preset;
	const client = connect();

	const presetRecord = await client.records.getOne('presets', id);
	const preset: Preset = {
		id: presetRecord.id,
		thumbnail: presetRecord.thumbnail,
		title: presetRecord.title,
		author: presetRecord.author,
		created: new Date(presetRecord.created),
		updated: new Date(presetRecord.updated)
	};

	const presetDataRecord = await client.records.getFullList('preset_data', undefined, {
		filter: `preset~'${id}'`
	});
	const presetData: Omit<PresetData, 'preset' | 'id'>[] = presetDataRecord.map((record) => ({
		data: record.data,
		title: record.title
	}));

	const authorName = (await client.records.getOne('profiles', preset.author)).name;

	return { preset, presetData, authorName };
}
