<script lang="ts">
	import 'normalize.css';
	import Header from '$lib/components/Header.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="layout">
	<header class="header">
		<Header />
	</header>

	<main class="content">
		<slot />
	</main>

	<nav class="sidebar">
		<Sidebar />
	</nav>

	<footer class="footer">
		<Footer />
	</footer>
</div>

<style lang="scss">
	.layout {
		display: grid;
		grid-template: auto / calc(100% - 400px) 400px;

		.header {
			grid-column: 1 / 3;
		}

		.content {
			grid-column: 1;
			padding-right: 35px;
			padding-left: 35px;
		}

		.sidebar {
			grid-column: 2;
			padding-right: 35px;
			padding-left: 35px;
		}

		.footer {
			grid-column: 1 / 3;
			padding: 35px;
		}
	}

	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
