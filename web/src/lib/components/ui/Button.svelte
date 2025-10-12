<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		children?: Snippet;
		loading?: boolean;
		loadingText?: string;
		class?: string;
		type?: 'button' | 'submit' | 'reset';
	}

	let { children, loading, loadingText, type, class: buttonClass }: Props = $props();
</script>

{#snippet content()}
	{#if loading}
		<span class="loading loading-spinner"></span>
		{#if loadingText}
			{loadingText}
		{/if}
	{:else}
		{@render children?.()}
	{/if}
{/snippet}

<button {type} disabled={loading} class={`btn ${buttonClass}`}>
	{@render content()}
</button>
