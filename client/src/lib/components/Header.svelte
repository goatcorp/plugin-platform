<script lang="ts">
	import 'normalize.css';
	import { connectBackend } from '$lib/backend';
	import type { User } from 'pocketbase';
	import { onMount } from 'svelte';

	let user: User | null = null;

	const loadCurrentUser = async () => {
		const backend = connectBackend();
		const model = backend.getCurrentUser();
		if (model && 'id' in model) {
			// The stored user data doesn't include the profile data
			user = await backend.app.users.getOne(model.id);
			return;
		}
	};

	onMount(async () => {
		await loadCurrentUser();
	});
</script>

<nav>
	<div><a href="/">Dalamud Plugin Presets</a></div>
	<div class="controls">
		<div class="search">
			<div>
				<form method="get" action="/search">
					<input type="search" placeholder="Search for presets..." name="q" />
					<button>Search</button>
				</form>
			</div>
		</div>
		{#if user == null}
			<a href="/login">Log in</a>
		{:else}
			<a href="/logout">Log out</a>
			<a href="/create">Create</a>
			<a href="/settings">Settings</a>
			<a href={`/user/${user.profile?.id}`}>{user.profile?.name}</a>
		{/if}
	</div>
</nav>

<style lang="scss">
	nav {
		padding: 35px;
		display: flex;
		justify-content: space-between;

		.controls {
			display: flex;
			justify-content: right;

			> * {
				margin-left: 8px;
			}
		}

		.search {
			display: flex;
		}
	}
</style>
