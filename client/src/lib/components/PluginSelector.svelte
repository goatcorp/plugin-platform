<script lang="ts">
	import type { Plugin } from '$lib/plugins';

	export let plugins: Plugin[];
	export let selectedPlugin: Plugin | null | undefined;

	export let readOnly: boolean = false;
	export let onSet: (plugin: Plugin | null | undefined) => void;
</script>

{#if !readOnly}
	<select>
		<option value="" disabled selected={selectedPlugin == null}>--Select a plugin--</option>
		{#each plugins as plugin}
			<option
				value={plugin.id}
				on:click={() => onSet(plugin)}
				selected={selectedPlugin?.id === plugin.id}>{plugin.name}</option
			>
		{/each}
	</select>
{:else}
	<span>{selectedPlugin?.name || ''}</span>
{/if}
