<script lang="ts">
	import type { Snippet } from 'svelte';
	
	interface Props {
		title?: string;
		children?: Snippet;
	}

	let { children, title }: Props = $props();
</script>

<div class="flex min-h-screen flex-col items-center justify-center bg-base-200 relative">
	<div class="w-full max-w-xl">
		<div class="card bg-base-100 shadow-xl">
			<div class="card-body p-10">
				<div class="flex flex-col items-center">
					<span class="mb-5 text-3xl font-bold text-primary">what it doo</span>
					<h2 class="mt-4 text-2xl font-bold">{title}</h2>
				</div>

				{@render children?.()}
			</div>
		</div>
	</div>

	<p class="absolute bottom-4 w-full text-center opacity-25">v1.0.0</p>
</div>
