<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;

	const origin = typeof location === 'undefined' ? '' : location.origin;
	const redirectUrl = `${origin}/oauth2/redirect`;

	const providerNames: Record<string, string> = {
		discord: 'Discord'
	};
</script>

<div>
	<span>Log in with...</span>
	<ul>
		{#each data.authProviders as provider}
			<li>
				<a
					href={`${provider.authUrl}${redirectUrl}`}
					class="link link-primary"
					on:click={() => {
						sessionStorage.setItem(
							'oauth2',
							JSON.stringify({
								name: provider.name,
								redirectUrl,
								state: provider.state,
								codeVerifier: provider.codeVerifier
							})
						);
					}}>{providerNames[provider.name]}</a
				>
			</li>
		{/each}
	</ul>
</div>
