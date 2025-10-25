<script lang="ts">
	import { openWebSocketConnection } from '$lib/stores/websocket.svelte';
	import '../app.css';

	let { children } = $props();

	openWebSocketConnection();	
</script>

{@render children?.()}

