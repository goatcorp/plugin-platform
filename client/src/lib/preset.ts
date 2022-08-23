import type { Record as PocketBaseRecord } from 'pocketbase';

export class Preset {
	id: string;
	thumbnail: string;
	title: string;
	author: string;
	spoiler: boolean;
	created: Date;
	updated: Date;

	private constructor(o: Record<string, unknown>) {
		this.id = o['id'] as string;
		this.thumbnail = o['thumbnail'] as string;
		this.title = o['title'] as string;
		this.author = o['author'] as string;
		this.spoiler = o['spoiler'] as boolean;
		this.created = o['created'] as Date;
		this.updated = o['updated'] as Date;
	}

	static fromRecord(r: PocketBaseRecord) {
		return new Preset({ ...r, created: new Date(r.created), updated: new Date(r.updated) });
	}
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
