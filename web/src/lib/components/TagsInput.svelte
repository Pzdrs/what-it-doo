<script lang="ts">
	let allOptions = ['JavaScript', 'TypeScript', 'Svelte', 'React', 'Vue'];
	let query = $state('');
	let tags: string[] = $state([]);

	function addTag(tag: string) {
        console.log('Adding tag:', tag);
		if (!tags.includes(tag)) {
			tags = [...tags, tag];
		}
		query = '';
	}

	function removeTag(tag: string) {
		tags = tags.filter((t) => t !== tag);
	}

	let focused = $state(false);
    let filtered = $derived(() => {
        const lowerQuery = query.toLowerCase();
        return allOptions.filter(
            (option) =>
                option.toLowerCase().includes(lowerQuery) && !tags.includes(option)
        );
    });
</script>

<div class="w-full max-w-md">
	<!-- Selected tags -->
	<div class="mb-2 flex flex-wrap gap-2">
		{#each tags as tag}
			<div class="badge badge-primary gap-1">
				{tag}
				<button type="button" class="ml-1 hover:text-white" onclick={() => removeTag(tag)}>
					✕
				</button>
			</div>
		{/each}
	</div>

	<!-- Input + dropdown -->
	<div class="relative">
		<input
			type="text"
			placeholder="Type or select a tag"
			bind:value={query}
			onfocus={() => (focused = true)}
			onblur={() => (focused = false)}
			class="input input-bordered w-full"
		/>

		{#if focused}
			<ul class="menu bg-base-100 rounded-box absolute z-10 mt-1 w-full shadow">
				{#each filtered() as option}
					<li>
						<a href="#curak" aria-label="Add tag" onclick={(e) => { e.preventDefault(); addTag(option); }}>{option}</a>
					</li>
				{/each}
			</ul>
		{/if}
	</div>
</div>
