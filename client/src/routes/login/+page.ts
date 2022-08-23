import { connectBackend } from '$lib/backend';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const backend = connectBackend();
	const providers = await backend.app.users.listAuthMethods();
	return { authProviders: providers.authProviders };
};
