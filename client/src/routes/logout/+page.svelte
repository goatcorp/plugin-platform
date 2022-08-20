<script lang="ts">
	import PocketBase from 'pocketbase';
	import { goto } from '$app/navigation';
	import { id } from '$lib/session';
	import { onMount } from 'svelte';

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const logout = async () => {
		const client = connect();
		client.authStore.clear();
		$id = null;
		try {
			await goto('/');
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
