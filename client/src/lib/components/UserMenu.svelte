<script lang="ts">
	import { UserIcon } from 'svelte-feather-icons';

	export let id: string;
	export let username: string;

	let open = false;
</script>

<div>
	<button
		on:click={() => {
			open = !open;
		}}
	>
		<UserIcon />
	</button>

	<div class="mount-point">
		{#if open}
			<div class="user-menu">
				<a href={`/user/${id}`}>{username}</a>
				<a href="/settings">Settings</a>
				<a href="/logout">Log out</a>
			</div>
		{/if}
	</div>
</div>

<style lang="scss">
	.mount-point {
		position: relative;
	}

	.user-menu {
		position: absolute;

		right: 0;

		padding: 4px;

		border: 1px solid black;
		border-radius: 5px;

		a {
			display: block;
		}
	}
</style>
