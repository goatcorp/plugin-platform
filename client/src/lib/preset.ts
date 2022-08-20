export interface Preset {
	id: string;
	thumbnail: string | null;
	title: string;
	author: string;
	created: Date;
	updated: Date;
}

export interface PresetData {
	id: string;
	preset: string;
	data: string;
	title: string;
}

export interface PresetStats {
	views: number; // Is it a good idea to give people a number to obsess over?
}
