<script>
	import { goto } from '$app/navigation';
	import { logout } from '$lib/api/client';
	import { AppRoute } from '$lib/constants';
	import ServerHealth from './ServerHealth.svelte';
	import ThemeSwitch from './ThemeSwitch.svelte';
	let menuOpen = false;

	const _logout = async () => {
		const { redirect_url } = await logout();
		await goto(redirect_url ?? AppRoute.AUTH_LOGIN);
	};
</script>

<div class="my-1 navbar h-20 bg-base-100 shadow-sm">
	<!-- LEVÁ STRANA -->
	<div class="navbar-start">
		<a href="/" class="btn text-xl text-primary btn-ghost">what it doo</a>
	</div>

	<!-- PRAVÁ STRANA – DESKTOP -->
	<div class="navbar-end hidden gap-2 lg:flex">
		<ServerHealth />
		<ThemeSwitch />

		<div class="divider divider-horizontal"></div>

		<!-- Notifikace -->
		<div class="dropdown dropdown-end">
			<button aria-label="Notifications" class="btn btn-circle btn-ghost">
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
					<span class="indicator-item badge badge-xs badge-primary"></span>
				</div>
			</button>
			<ul class="dropdown-content menu z-10 w-52 rounded-box bg-base-100 p-2 shadow-sm">
				<li><a>Item 1</a></li>
				<li><a>Item 2</a></li>
			</ul>
		</div>

		<!-- Avatar -->
		<div class="dropdown dropdown-end">
			<div tabindex="0" role="button" class="btn avatar btn-circle btn-ghost">
				<div class="w-10 rounded-full">
					<img
						alt="user avatar"
						src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
					/>
				</div>
			</div>
			<ul class="dropdown-content menu z-10 mt-3 w-52 menu-sm rounded-box bg-base-100 p-2 shadow">
				<li>
					<a href="#profile" class="justify-between">Profile <span class="badge">New</span></a>
				</li>
				<li><a href="#settings">Settings</a></li>
				<li><button onclick={_logout}>Logout</button></li>
			</ul>
		</div>
	</div>

	<!-- MOBILE BURGER -->
	<div class="navbar-end lg:hidden">
		<button class="btn btn-circle btn-ghost" onclick={() => (menuOpen = !menuOpen)}>
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
	<div class="border-t border-base-200 bg-base-100 shadow-md lg:hidden">
		<div class="flex flex-col items-center gap-3 p-4">
			<ServerHealth />
			<ThemeSwitch />
			<a href="#notifications" class="btn justify-start btn-ghost">Notifications</a>
			<a href="#profile" class="btn justify-start btn-ghost">Profile</a>
			<a href="#settings" class="btn justify-start btn-ghost">Settings</a>
			<a href="#logout" class="btn justify-start btn-ghost">Logout</a>
		</div>
	</div>
{/if}
