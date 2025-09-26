<script lang="ts">
	// --- reactive state (Svelte 5 runes) ---
	let allOptions = ['JavaScript', 'TypeScript', 'Svelte', 'React', 'Vue'];

	let query = $state(''); // text in the input
	let tags: string[] = $state([]); // chosen tags
	let focused = $state(true); // whether the input is focused

	// derived list filtered by query and not already chosen
	let filtered = $derived(() => {
		const lower = query.toLowerCase();
		return allOptions.filter((o) => o.toLowerCase().includes(lower) && !tags.includes(o));
	});

	function addTag(tag: string) {
		console.log('Adding tag:', tag);
		if (!tags.includes(tag)) tags = [...tags, tag];
		query = '';
	}

	function removeTag(tag: string) {
		tags = tags.filter((t) => t !== tag);
	}
</script>

<div class="w-full">
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
			onblur={() => setTimeout(() => (focused = false), 100)}
			class="input input-bordered w-full"
		/>

		{#if focused && filtered().length}
			<ul class="menu bg-base-100 rounded-box absolute z-10 mt-1 w-full shadow">
				{#each filtered() as option}
					<li>
						<div class="flex gap-3">
							<div class="avatar">
								<div class="w-12 rounded-full">
									<img
										src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p200x200_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=bW-Nk_SfhrEQ7kNvwEYj1ab&_nc_oc=AdmnK9O53ElfigOxXct-Vi8G0jm10Q64AR71Rb62wFxLKOt4gJJsq8UQPuRQm9IpmMI&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=iTU3jBr-XOTG0tPi5IzqBA&oh=00_AfZm3SJPNquCb3o8xQWx0bJc7fl10-3xZwVbUlw_WVim9g&oe=68D61548"
										alt="Chat Icon"
									/>
								</div>
							</div>
							<button type="button" class="w-full text-left" onclick={() => addTag(option)}>
								{option}
							</button>
						</div>
					</li>
				{/each}
			</ul>
		{/if}
	</div>
</div>
