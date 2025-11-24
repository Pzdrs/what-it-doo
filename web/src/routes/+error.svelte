<script>
	import { page } from '$app/state';
	import { AppRoute } from '$lib/constants';
	import { t } from 'svelte-i18n';

</script>

<main class="bg-base-200 grid min-h-screen place-items-center px-6 py-24 sm:py-32 lg:px-8">
	<div class="text-center">
		<p class="text-primary text-base font-semibold">{page.error?.status || '404'}</p>

		<h1 class="mt-4 text-5xl font-bold tracking-tight sm:text-7xl">{page.error?.message}</h1>

		<p class="text-base-content/70 mt-6 text-lg sm:text-xl">
			{$t('errors.unknown_error')}
		</p>

		<div class="mt-10 flex items-center justify-center gap-6">
			<a href="{AppRoute.HOME}" class="btn btn-primary"> {$t('go_back_home')} </a>
		</div>
	</div>
</main>
