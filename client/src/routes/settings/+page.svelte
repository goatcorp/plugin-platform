<script lang="ts">
	import { onMount } from 'svelte';
	import { id } from '$lib/session';
	import { getSettings, updateSettings, type Settings } from '$lib/settings';

	let settings: Settings | null = null;
	let updated: Partial<Settings> = {};

	type UpdateSetting<T extends Settings = Settings, U extends keyof T = keyof T> = (
		key: keyof T,
		value: T[U]
	) => void;

	const update: UpdateSetting = (key, value) => {
		if (settings == null) {
			return;
		}

		settings[key] = value;
		updated[key] = value;
		console.log(updated);
	};

	const save = async () => {
		if (settings == null) {
			return;
		}

		await updateSettings(updated);
	};

	onMount(async () => {
		settings = await getSettings($id);
	});
</script>

<h1>Settings</h1>

<div class="gap" />

<form
	on:submit={(e) => {
		e.preventDefault();
		save();
	}}
>
	<label for="show_spoilers">Automatically show spoilers</label>
	<input
		type="checkbox"
		name="show_spoilers"
		checked={settings?.show_spoilers}
		on:change={(e) => update('show_spoilers', e.currentTarget.checked)}
	/>
	<br />
	<button>Save</button>
</form>
