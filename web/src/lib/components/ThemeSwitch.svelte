<script lang="ts">
	const states = {
		dark: true,
		light: false
	};
	let currentTheme = 'light';

	function setTheme(theme: string) {
		currentTheme = theme;
		document.documentElement.setAttribute('data-theme', theme);
		localStorage.setItem('daisy-theme', theme);
		document.querySelector("input[name='theme']")!.checked = states[theme];
	}

	$effect(() => {
		const savedTheme = localStorage.getItem('daisy-theme');
		if (savedTheme) {
			setTheme(savedTheme);
		} else {
			setTheme('light');
		}
	});

	function toggleTheme() {
		const nextTheme = currentTheme === 'light' ? 'dark' : 'light';
		setTheme(nextTheme);
	}
</script>

<label class="flex cursor-pointer gap-2">
	<svg
		xmlns="http://www.w3.org/2000/svg"
		width="20"
		height="20"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		stroke-width="2"
		stroke-linecap="round"
		stroke-linejoin="round"
	>
		<circle cx="12" cy="12" r="5" />
		<path
			d="M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4"
		/>
	</svg>
	<input
		type="checkbox"
		name="theme"
		value="synthwave"
		class="toggle theme-controller"
		onclick={() => toggleTheme()}
	/>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		width="20"
		height="20"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		stroke-width="2"
		stroke-linecap="round"
		stroke-linejoin="round"
	>
		<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
	</svg>
</label>
