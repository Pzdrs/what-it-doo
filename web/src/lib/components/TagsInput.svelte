<script lang="ts">
	// --- reactive state (Svelte 5 runes) ---
	let allOptions = ['JavaScript', 'TypeScript', 'Svelte', 'React', 'Vue'];

	let query = $state(''); // text in the input
	let tags: string[] = $state([]); // chosen tags
	let focused = $state(false); // whether the input is focused

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
			<div class="bg-base-200 rounded-field text-base-content gap-1 p-2">
				<div class="avatar mr-2">
					<div class="w-12 rounded-full">
						<img
							src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p100x100_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=KBdDtGYe2KEQ7kNvwFDpOhW&_nc_oc=Adlz9gFlsgmnZdW0y8Hk2W2uryRA3duwN5KZf5g4bX_RH-RnCC_jQEIh_6ns4sQBjIo&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=_yd9s4D9R7GqfZ-DtwCVWQ&oh=00_Afar8I8NEQnekKKjagUfvYNsvrOyjnKez2h0w7dM6F1kXg&oe=68DC7488"
							alt="Chat Icon"
						/>
					</div>
				</div>
				{tag}
				<button
					type="button"
					class="hover:text-base-content/50 ml-1"
					onclick={() => removeTag(tag)}
				>
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
						<button type="button" class="flex gap-3" onclick={() => addTag(option)}>
							<div class="avatar">
								<div class="w-12 rounded-full">
									<img
										src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p100x100_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=KBdDtGYe2KEQ7kNvwFDpOhW&_nc_oc=Adlz9gFlsgmnZdW0y8Hk2W2uryRA3duwN5KZf5g4bX_RH-RnCC_jQEIh_6ns4sQBjIo&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=_yd9s4D9R7GqfZ-DtwCVWQ&oh=00_Afar8I8NEQnekKKjagUfvYNsvrOyjnKez2h0w7dM6F1kXg&oe=68DC7488"
										alt="Chat Icon"
									/>
								</div>
							</div>
							{option}
						</button>
					</li>
				{/each}
			</ul>
		{/if}
	</div>
</div>
