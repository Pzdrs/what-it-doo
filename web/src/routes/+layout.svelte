<script lang="ts">
	import { page } from '$app/state';
	import { openWebSocketConnection } from '$lib/stores/websocket.svelte';
	import '../app.css';

	let { children } = $props();

	openWebSocketConnection();
</script>

<svelte:head>
	<title>{page.data.meta?.title || 'Web'} - What It Doo</title>
</svelte:head>

{@render children?.()}
