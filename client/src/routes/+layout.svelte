<script lang="ts">
	import '../app.css';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="container mx-auto w-5/6">
	<header class="p-9">
		<Header />
	</header>

	<main class="p-9">
		<slot />
	</main>

	<footer class="p-9">
		<Footer />
	</footer>
</div>
