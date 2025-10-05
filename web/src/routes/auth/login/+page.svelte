<script lang="ts">
	import { goto } from '$app/navigation';
	import { login } from '$lib/api/client';
	import AuthPageLayout from '$lib/components/layout/AuthPageLayout.svelte';
	import { t } from 'svelte-i18n';
	import type { PageData } from './$types';
	import { getTranslatedError } from '$lib/utils/handle-error';
	import { setUser } from '$lib/stores/user.svelte';

	interface Props {
		data: PageData;
	}

	let { data }: Props = $props();

	let errorMessage: string = $state('');
	let email = $state('');
	let password = $state('');
	let remember_me = $state(false);
	let loading = $state(false);

	const onSuccess = async (user) => {
		setUser(user);
		await goto(data.continueUrl, { invalidateAll: true });
	};

	const handleLogin = async () => {
		try {
			errorMessage = '';
			loading = true;
			const user = await login({ email, password, remember_me });
			await onSuccess(user);
		} catch (error) {
			errorMessage = getTranslatedError($t, error);
			password = '';
			loading = false;
		}
	};
	const handleIdpLogin = async () => {};

	const onsubmit = async (event: Event) => {
		event.preventDefault();
		await handleLogin();
	};
</script>

<AuthPageLayout title={$t('sign_in_please')}>
	<form {onsubmit} class="mt-6 space-y-4">
		{#if errorMessage}
			<div role="alert" class="mb-10 alert alert-error">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-6 w-6 shrink-0 stroke-current"
					fill="none"
					viewBox="0 0 24 24"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
					/>
				</svg>
				<span>{errorMessage}</span>
			</div>
		{/if}

		<div>
			<label class="label">
				<span class="label-text">{$t('email_address')}</span>
			</label>
			<input
				type="email"
				placeholder={$t('email_placeholder')}
				bind:value={email}
				class="input-bordered input w-full"
				required
			/>
		</div>

		<div>
			<label class="label">
				<span class="label-text">{$t('password')}</span>
			</label>
			<input
				type="password"
				placeholder="••••••••"
				bind:value={password}
				class="input-bordered input w-full"
				required
			/>
		</div>

		<div class="flex items-center justify-between">
			<label class="label cursor-pointer gap-2">
				<input type="checkbox" bind:checked={remember_me} class="checkbox checkbox-primary" />
				<span class="label-text">{$t('remember_me')}</span>
			</label>
			<a href="#" class="link text-sm link-primary">{$t('forgot_password')}</a>
		</div>

		<button type="submit" class="btn w-full btn-primary">{$t('sign_in')}</button>

		<div class="divider">{$t('or_continue_with')}</div>

		<div class="flex flex-wrap justify-center gap-3">
			<button onclick={handleIdpLogin} type="button" class="btn w-full gap-2 btn-outline">
				<img
					src="https://www.svgrepo.com/show/303108/google-icon-logo.svg"
					alt="Google"
					class="h-5 w-5"
				/>
				Google
			</button>
		</div>
	</form>

	<p class="mt-6 text-center text-sm">
		{$t('no_account')}
		<a href="/auth/register" class="link link-primary">{$t('sign_up')}</a>
	</p>
</AuthPageLayout>
