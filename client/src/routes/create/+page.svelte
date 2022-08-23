<script lang="ts">
	import { goto } from '$app/navigation';
	import { connectBackend } from '$lib/backend';

	const createPreset = async (data: FormData) => {
		const backend = connectBackend();

		const user = backend.app.authStore.model;
		if (!('id' in user && 'profile' in user)) {
			return;
		}

		data.set('author', user.profile?.id || '');

		const preset = new FormData();
		if (data.has('title')) preset.set('title', data.get('title')!);
		if (data.has('thumbnail')) preset.set('thumbnail', data.get('thumbnail')!);
		if (data.has('author')) preset.set('author', data.get('author')!);

		try {
			const presetRecord = await backend.app.records.create('presets', preset);

			const presetData = new FormData();
			if (data.has('data0_title')) presetData.set('title', data.get('data0_title')!);
			if (data.has('data0_data')) presetData.set('data', data.get('data0_data')!);
			presetData.set('preset', presetRecord.id);

			try {
				await backend.app.records.create('preset_data', presetData);
			} catch (err) {
				console.error(err);
			}

			try {
				await goto(`/preset/${presetRecord.id}`);
			} catch (err) {
				console.error(err);
			}
		} catch (err) {
			console.error(err);
		}
	};
</script>

<h1>Create a preset</h1>

<form
	on:submit={(e) => {
		e.preventDefault();
		const form = new FormData(e.currentTarget);
		createPreset(form);
	}}
>
	<label for="title">Title</label>
	<input type="text" name="title" />
	<br />
	<label for="thumbnail">Thumbnail</label>
	<input type="file" name="thumbnail" />
	<br />
	<label for="data0_title">Data title</label>
	<input type="text" name="data0_title" />
	<br />
	<label for="data0_data">Data</label>
	<input type="text" name="data0_data" />
	<br />
	<button>Create</button>
</form>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
