import type { AppUser } from './app-user';

export interface Preset {
	id: string;
	thumbnail: string | null;
	title: string;
	author: AppUser;
}
