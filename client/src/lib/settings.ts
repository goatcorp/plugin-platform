import PocketBase from 'pocketbase';

export interface Settings {
	show_spoilers: boolean;
}

export async function getSettings(userId?: string | null | undefined): Promise<Settings> {
	const client = new PocketBase('http://127.0.0.1:8090');
	const settings = await client.records.getFullList('user_settings');
	if (settings.length === 0) {
		const newSettings: Settings = {
			show_spoilers: false
		};

		try {
			await client.records.create('user_settings', { ...newSettings, user: userId });
		} catch (err) {
			console.error(err);
		}

		return newSettings;
	}

	return {
		show_spoilers: settings[0].show_spoilers
	};
}

export async function updateSettings(settings: Partial<Settings>): Promise<void> {
	const client = new PocketBase('http://127.0.0.1:8090');
	const existing = (await client.records.getFullList('user_settings'))[0];
	try {
		await client.records.update('user_settings', existing.id, settings);
	} catch (err) {
		console.error(err);
	}
}
