<script>
	import ServerHealth from './ServerHealth.svelte';
	import ThemeSwitch from './ThemeSwitch.svelte';
	let menuOpen = false;
</script>

<div class="navbar bg-base-100 my-1 h-20 shadow-sm">
	<!-- LEVÁ STRANA -->
	<div class="navbar-start">
		<a href="/" class="btn btn-ghost text-primary text-xl">what it doo</a>
	</div>

	<!-- PRAVÁ STRANA – DESKTOP -->
	<div class="navbar-end hidden gap-2 lg:flex">
		<ServerHealth />
		<ThemeSwitch />

		<div class="divider divider-horizontal"></div>

		<!-- Notifikace -->
		<div class="dropdown dropdown-end">
			<button aria-label="Notifications" class="btn btn-ghost btn-circle">
				<div class="indicator">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
						/>
					</svg>
					<span class="badge badge-xs badge-primary indicator-item"></span>
				</div>
			</button>
			<ul class="dropdown-content menu bg-base-100 rounded-box z-10 w-52 p-2 shadow-sm">
				<li><a>Item 1</a></li>
				<li><a>Item 2</a></li>
			</ul>
		</div>

		<!-- Avatar -->
		<div class="dropdown dropdown-end">
			<div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
				<div class="w-10 rounded-full">
					<img
						alt="user avatar"
						src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
					/>
				</div>
			</div>
			<ul class="menu menu-sm dropdown-content bg-base-100 rounded-box z-10 mt-3 w-52 p-2 shadow">
				<li>
					<a href="#profile" class="justify-between">Profile <span class="badge">New</span></a>
				</li>
				<li><a href="#settings">Settings</a></li>
				<li><a href="#logout">Logout</a></li>
			</ul>
		</div>
	</div>

	<!-- MOBILE BURGER -->
	<div class="navbar-end lg:hidden">
		<button class="btn btn-ghost btn-circle" on:click={() => (menuOpen = !menuOpen)}>
			<!-- ikonka hamburgeru -->
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-6 w-6"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 6h16M4 12h16M4 18h16"
				/>
			</svg>
		</button>
	</div>
</div>

<!-- MOBILE MENU -->
{#if menuOpen}
	<div class="bg-base-100 border-base-200 border-t shadow-md lg:hidden">
		<div class="flex flex-col items-center gap-3 p-4">
			<ServerHealth />
			<ThemeSwitch />
			<a href="#notifications" class="btn btn-ghost justify-start">Notifications</a>
			<a href="#profile" class="btn btn-ghost justify-start">Profile</a>
			<a href="#settings" class="btn btn-ghost justify-start">Settings</a>
			<a href="#logout" class="btn btn-ghost justify-start">Logout</a>
		</div>
	</div>
{/if}
