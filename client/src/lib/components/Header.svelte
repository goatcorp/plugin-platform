<script lang="ts">
	import 'normalize.css';
	import { connectBackend } from '$lib/backend';
	import type { User } from 'pocketbase';
	import { onMount } from 'svelte';
	import UserMenu from './UserMenu.svelte';
	import { SearchIcon } from 'svelte-feather-icons';

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
					<button><SearchIcon /></button>
				</form>
			</div>
		</div>
		{#if user?.profile == null}
			<a href="/login">Log in</a>
		{:else}
			<a href="/create">Create</a>
			<UserMenu id={user.profile.id} />
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
				margin-right: 4px;
				margin-left: 4px;
			}
		}

		.search {
			display: flex;
		}
	}
</style>
