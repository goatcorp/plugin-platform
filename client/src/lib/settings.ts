import PocketBase from 'pocketbase';
import { connectBackend } from './backend';

export interface Settings {
	show_spoilers: boolean;
}

export async function getSettings(userId?: string | null | undefined): Promise<Settings> {
	const backend = connectBackend();
	const settings = await backend.app.records.getFullList('user_settings');
	if (settings.length === 0) {
		const newSettings: Settings = {
			show_spoilers: false
		};

		if (userId != null) {
			try {
				await backend.app.records.create('user_settings', { ...newSettings, user: userId });
			} catch (err) {
				console.error(err);
			}
		}

		return newSettings;
	}

	return {
		show_spoilers: settings[0].show_spoilers
	};
}

export async function updateSettings(settings: Partial<Settings>): Promise<void> {
	const backend = connectBackend();
	const existing = (await backend.app.records.getFullList('user_settings'))[0];
	try {
		await backend.app.records.update('user_settings', existing.id, settings);
	} catch (err) {
		console.error(err);
	}
}
