import type { Preset, PresetStats } from '$lib/preset';
import type { Profile } from '$lib/profile';
import PocketBase from 'pocketbase';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export async function load({ params }: { params: Record<string, string> }) {
	const id = params.user;
	const client = connect();

	const profileRecord = await client.records.getOne('profiles', id);
	const profile: Profile = {
		name: profileRecord.name,
		avatar: profileRecord.avatar
	};

	const presetRecords = await client.records.getFullList('presets', undefined, {
		filter: `author='${id}'`
	});
	const presets: Preset[] = presetRecords.map((record) => ({
		id: record.id,
		thumbnail: record.thumbnail,
		title: record.title,
		author: record.author,
		views: record.views,
		created: new Date(record.created),
		updated: new Date(record.updated)
	}));

	const presetStats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		presetStats[preset.id] = presetStatsRecords
			.map((record) => ({
				views: record.views
			}))
			.reduce((agg, next) => {
				agg.views += next.views;
				return agg;
			});
	}

	return { profile, presets, presetStats };
}
