import type { Preset } from '$lib/preset';
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
		created: new Date(record.created),
		updated: new Date(record.updated)
	}));

	return { profile, presets };
}
