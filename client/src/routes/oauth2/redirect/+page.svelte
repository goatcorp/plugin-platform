<script lang="ts">
	import PocketBase from 'pocketbase';
	import { onMount } from 'svelte';

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const exchangeToken = async () => {
		const dataRaw = sessionStorage.getItem('oauth2');
		if (dataRaw == null) {
			console.error('oauth2 data not set');
			return;
		}

		const data = JSON.parse(dataRaw);
		const url = new URL(location.toString());

		const providerCode = url.searchParams.get('code');
		if (providerCode == null) {
			console.error('code not present in URL parameters');
			return;
		}

		const providerState = url.searchParams.get('state');
		if (providerState == null) {
			console.error('state not present in URL parameters');
			return;
		}

		if (data.state !== providerState) {
			console.error(`CSRF state mismatch! Client: "${data.state}"; Provider: "${providerState}"`);
			return;
		}

		const client = connect();
		try {
			await client.users.authViaOAuth2(
				data.name,
				providerCode,
				data.codeVerifier,
				data.redirectUrl
			);

			// Reloads state
			location.assign('/');
		} catch (err) {
			console.error(err);
		}
	};

	onMount(async () => {
		await exchangeToken();
	});
</script>

<pre>Redirecting...</pre>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
