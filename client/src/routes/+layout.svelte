<script lang="ts">
	import 'normalize.css';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="layout">
	<header class="header">
		<Header />
	</header>

	<main class="content">
		<slot />
	</main>

	<footer class="footer">
		<Footer />
	</footer>
</div>

<style lang="scss">
	.layout {
		width: 70%;
		margin: auto;

		.content {
			padding-right: 35px;
			padding-left: 35px;
		}

		.footer {
			padding: 35px;
		}
	}

	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
