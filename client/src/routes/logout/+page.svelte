<script lang="ts">
	import { connectBackend } from '$lib/backend';

	import { onMount } from 'svelte';

	const logout = async () => {
		const backend = connectBackend();
		backend.app.authStore.clear();
		try {
			// Reloads state
			location.assign('/');
		} catch (err) {
			console.error(err);
		}
	};

	onMount(async () => {
		await logout();
	});
</script>

<p>Redirecting...</p>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
