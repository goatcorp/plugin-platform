<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { connectBackend } from '$lib/backend';

	export let data: PageData;

	const updatePreset = async (form: FormData) => {
		const backend = connectBackend();

		const preset = new FormData();
		if (form.has('title')) preset.set('title', form.get('title')!);
		if (form.has('thumbnail')) preset.set('thumbnail', form.get('thumbnail')!);

		try {
			await backend.app.records.update('presets', data.preset.id, preset);

			const presetData = new FormData();
			if (form.has('data0_title')) presetData.set('title', form.get('data0_title')!);
			if (form.has('data0_data')) presetData.set('data', form.get('data0_data')!);

			try {
				await backend.app.records.update('preset_data', data.presetData[0].id, presetData);
			} catch (err) {
				console.error(err);
			}

			// TODO: the other preset data entries
			try {
				await goto(`/preset/${data.preset.id}`);
			} catch (err) {
				console.error(err);
			}
		} catch (err) {
			console.error(err);
		}
	};
</script>

<h1>Editing {data.preset.title}</h1>

<form
	on:submit={(e) => {
		e.preventDefault();
		const form = new FormData(e.currentTarget);
		updatePreset(form);
	}}
>
	<label for="title">Title</label>
	<input type="text" name="title" value={data.preset.title} />
	<br />
	<label for="thumbnail">Thumbnail</label>
	<input type="file" name="thumbnail" />
	<br />
	<label for="data0_title">Data title</label>
	<input type="text" name="data0_title" value={data.presetData[0].title} />
	<br />
	<label for="data0_data">Data</label>
	<input type="text" name="data0_data" value={data.presetData[0].data} />
	<br />
	<button>Save</button>
</form>
