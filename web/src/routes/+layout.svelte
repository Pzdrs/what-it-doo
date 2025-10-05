<script lang="ts">
	import { logout } from '$lib/api/client';
	import '../app.css';

	let { children } = $props();

	async function onclick() {
		await logout();
	}
</script>

{@render children?.()}

