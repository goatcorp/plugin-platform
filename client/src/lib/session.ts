import { writable, type Writable } from 'svelte/store';

export const id: Writable<string | null> = writable(
	typeof window !== 'undefined' ? localStorage.getItem('id') : null || null
);
id.subscribe((x) => {
	if (typeof window === 'undefined') {
		return;
	}

	if (x != null) {
		localStorage.setItem('id', x);
	} else {
		localStorage.removeItem('id');
	}
});
