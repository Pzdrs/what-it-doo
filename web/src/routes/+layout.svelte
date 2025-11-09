<script lang="ts">
	import { page } from '$app/state';
	import '../app.css';

	let { children } = $props();
</script>

<svelte:head>
	<title>{page.data.meta?.title || 'Web'} - What It Doo</title>
</svelte:head>

{@render children?.()}
