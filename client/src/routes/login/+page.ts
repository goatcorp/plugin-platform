import PocketBase from 'pocketbase';
import type { PageLoad } from './$types';

const connect = () => {
	return new PocketBase('http://127.0.0.1:8090');
};

export const load: PageLoad = async () => {
	const client = connect();
	const providers = await client.users.listAuthMethods();
	return { authProviders: providers.authProviders };
};
