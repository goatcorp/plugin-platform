<script lang="ts">
	import { UserIcon } from 'svelte-feather-icons';

	export let id: string;

	let open = false;

	const toggleMenu = () => {
		open = !open;
	};

	const closeMenu = () => {
		open = false;
	};
</script>

<div>
	<button on:click={toggleMenu}>
		<UserIcon />
	</button>

	<div class="mount-point">
		{#if open}
			<div class="user-menu">
				<a href={`/user/${id}`} on:click={closeMenu}>My account</a>
				<a href="/settings" on:click={closeMenu}>Settings</a>
				<a href="/logout" on:click={closeMenu}>Log out</a>
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
