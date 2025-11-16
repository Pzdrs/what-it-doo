<script>
	import { goto } from '$app/navigation';
	import { logout } from '$lib/api/client';
	import user from '$lib/assets/user.png?enhanced';
	import { AppRoute } from '$lib/constants';
	import { userStore } from '$stores/user.svelte';
	import ThemeSwitch from './ThemeSwitch.svelte';
	let menuOpen = false;

	const _logout = async () => {
		const { redirect_url } = await logout();
		await goto(redirect_url ?? AppRoute.AUTH_LOGIN, { invalidateAll: true });
	};
</script>

<div class="navbar bg-base-100 my-1 h-20 shadow-sm">
	<!-- LEVÁ STRANA -->
	<div class="navbar-start">
		<a href="/" class="btn text-primary btn-ghost text-xl">what it doo</a>
	</div>

	<!-- PRAVÁ STRANA – DESKTOP -->
	<div class="navbar-end hidden gap-2 lg:flex">
		<ThemeSwitch />

		<div class="divider divider-horizontal"></div>

		<!-- Avatar -->
		{#if userStore.user}
			<div class="dropdown dropdown-end">
				<div tabindex="0" role="button" class="btn avatar btn-circle btn-ghost">
					<div class="w-10 rounded-full">
						{#if userStore.user.avatar_url}
							<img alt="user avatar" src={userStore.user.avatar_url} />
						{:else}
							<enhanced:img alt="default user avatar" src={user} />
						{/if}
					</div>
				</div>
				<ul class="dropdown-content menu menu-sm rounded-box bg-base-100 z-10 mt-3 w-52 p-2 shadow">
					<li><button onclick={_logout}>Logout</button></li>
				</ul>
			</div>
		{/if}
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
	<div class="border-base-200 bg-base-100 border-t shadow-md lg:hidden">
		<div class="flex flex-col items-center gap-3 p-4">
			<ThemeSwitch />
			<a href="#logout" class="btn btn-ghost justify-start">Logout</a>
		</div>
	</div>
{/if}
