export interface Preset {
	id: string;
	thumbnail: string | null;
	title: string;
	author: string;
	created: Date;
}

export interface PresetData {
	id: string;
	preset: string;
	data: string;
}
