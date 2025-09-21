<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<nav class="bg-gray-400 flex justify-end">
	<button>Dark Mode</button>
	<span>Profile</span>
</nav>

{@render children?.()}