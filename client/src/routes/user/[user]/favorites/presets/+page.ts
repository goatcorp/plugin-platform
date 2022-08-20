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
		id: profileRecord.id,
		name: profileRecord.name,
		avatar: profileRecord.avatar
	};

	const favorites = await client.records.getFullList('user_preset_favorites');
	const presets: Preset[] = [];
	for (const favorite of favorites) {
		const preset = await client.records.getOne('presets', favorite.preset);
		presets.push({
			id: preset.id,
			thumbnail: preset.thumbnail,
			title: preset.title,
			author: preset.author,
			created: new Date(preset.created),
			updated: new Date(preset.updated)
		});
	}

	const presetAuthors: Record<string, string> = {};
	for (const preset of presets) {
		const author = await client.records.getOne('profiles', preset.author);
		presetAuthors[preset.author] = author.name;
	}

	const presetStats: Record<string, PresetStats> = {};
	for (const preset of presets) {
		const presetStatsRecords = await client.records.getFullList('preset_stats', undefined, {
			filter: `preset~'${preset.id}'`
		});
		presetStats[preset.id] = presetStatsRecords
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

	return { profile, presets, presetAuthors, presetStats };
}
