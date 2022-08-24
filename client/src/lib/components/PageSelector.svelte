<script lang="ts">
	export let page: number;
	export let totalPages: number;
</script>

<div>
	{#if page > 1}
		<a href={`?page=${page - 1}`}>Previous</a>
	{:else}
		<a href={`?page=${page - 1}`} disabled tabindex="-1">Previous</a>
	{/if}
	<a href={`?page=${page}`} disabled tabindex="-1">{page}</a>
	{#if page < totalPages}
		<a href={`?page=${page + 1}`}>Next</a>
	{:else}
		<a href={`?page=${page + 1}`} disabled tabindex="-1">Next</a>
	{/if}
</div>

<style lang="scss">
	a[disabled] {
		pointer-events: none;
		text-decoration: none;
		color: inherit;
	}
</style>
