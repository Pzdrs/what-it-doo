<script lang="ts">
	import { goto } from '$app/navigation';
	import { login, type DtoUserDetails } from '$lib/api/client';
	import AuthPageLayout from '$lib/components/layout/AuthPageLayout.svelte';
	import Alert from '$lib/components/ui/Alert.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { userStore } from '$lib/stores/user.svelte';
	import { getTranslatedError } from '$lib/utils/handle-error';
	import { t } from 'svelte-i18n';
	import type { PageData } from './$types';

	interface Props {
		data: PageData;
	}

	let { data }: Props = $props();

	let errorMessage: string = $state('');
	let email = $state('');
	let password = $state('');
	let remember_me = $state(false);
	let loading = $state(false);

	const onSuccess = async (user: DtoUserDetails) => {
		userStore.user = user;	
		await goto(data.continueUrl, { invalidateAll: true });
	};

	const handleLogin = async () => {
		try {
			errorMessage = '';
			loading = true;
			const { user } = await login({ email, password, remember_me });
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
		<Alert type="error" message={errorMessage} />

		<div>
			<label for="email" class="label">
				<span class="label-text">{$t('email_address')}</span>
			</label>
			<input
				name="email"
				type="email"
				placeholder={$t('email_placeholder')}
				bind:value={email}
				class="input-bordered input w-full"
				required
			/>
		</div>

		<div>
			<label for="password" class="label">
				<span class="label-text">{$t('password')}</span>
			</label>
			<input
				name="password"
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

		<Button {loading} loadingText={$t('sign_in_loading')} type="submit" class="w-full btn-primary">
			{$t('sign_in')}
		</Button>
	</form>

	<p class="mt-6 text-center text-sm">
		{$t('no_account')}
		<a href="/auth/register" class="link link-primary">{$t('sign_up')}</a>
	</p>
</AuthPageLayout>
