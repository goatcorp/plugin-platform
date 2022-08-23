import type { Record } from 'pocketbase';

export interface ListResult<M extends Record> {
	page: number;
	perPage: number;
	totalItems: number;
	totalPages: number;
	items: Array<M>;
}
