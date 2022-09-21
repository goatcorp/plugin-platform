<script lang="ts">
	import type { Tag } from '$lib/tags';

	export let tagOptions: Tag[];
	export let tags: string[];
	export let readOnly: boolean = false;
	export let onSearch: (label: string) => void;
	export let onAdd: (id: string) => void;
	export let onRemove: (label: string) => void;
</script>

{#each tags as tag}
	<div>
		{#if !readOnly}
			<button on:click={() => onRemove(tag)}>Remove</button>
		{/if}
		<div class="badge badge-secondary">{tag}</div>
	</div>
{/each}
{#if tags.length === 0}
	<span>No tags.</span>
{/if}

{#if !readOnly}
	<div>
		<input type="text" on:keyup={(e) => onSearch(e.currentTarget.value)} />
		<select class="block">
			{#each tagOptions as tag}
				<option value={tag.id} on:click={() => onAdd(tag.id)}>{tag.label}</option>
			{/each}
		</select>
	</div>
{/if}
