<script lang="ts">
	export let page: number;
	export let totalPages: number;
</script>

<div>
	{#if page > 1}
		<a href={`?page=${page - 1}`} class="link link-primary">Previous</a>
	{:else}
		<a href={`?page=${page - 1}`} class="pointer-events-none" tabindex="-1">Previous</a>
	{/if}
	<a href={`?page=${page}`} class="pointer-events-none" tabindex="-1">{page}</a>
	{#if page < totalPages}
		<a href={`?page=${page + 1}`} class="link link-primary">Next</a>
	{:else}
		<a href={`?page=${page + 1}`} class="pointer-events-none" tabindex="-1">Next</a>
	{/if}
</div>
