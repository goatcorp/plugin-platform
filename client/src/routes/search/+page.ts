import PocketBase from 'pocketbase';
import type { Preset } from '$lib/preset';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export async function load({ url }: { url: URL }) {
	const query = url.searchParams.get('q') || '';
	const client = connect();
	const records = await client.records.getList('presets', 1, undefined, {
		filter: `title~'${query}'`
	});
	const results: Preset[] = records.items.map((item) => ({
		id: item.id,
		thumbnail: item.thumbnail,
		title: item.title,
		author: item.author,
		created: new Date(item.created),
		updated: new Date(item.updated)
	}));

	const authors: Record<string, string> = {};
	for (const result of results) {
		const user = await client.records.getOne('profiles', result.author);
		authors[result.author] = user.name;
	}

	return { results, authors, query };
}
