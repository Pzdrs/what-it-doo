<script lang="ts">
	import { goto } from '$app/navigation';
	import { register, type DtoUserDetails } from '$lib/api/client';
	import AuthPageLayout from '$lib/components/layout/AuthPageLayout.svelte';
	import Alert from '$lib/components/ui/Alert.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { AppRoute } from '$lib/constants';
	import { setUser } from '$lib/stores/user.svelte';
	import { getTranslatedError, isProblemType, toProblemDetails } from '$lib/utils/handle-error';
	import { t } from 'svelte-i18n';

	const passwordMinLength = 8;

	let errorMessage: string = $state('');
	let email = $state('');
	let name = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let loading = $state(false);

	const onSuccess = async (user: DtoUserDetails) => {
		setUser(user);
		await goto(AppRoute.HOME, { invalidateAll: true });
	};

	const handleRegister = async () => {
		try {
			errorMessage = '';
			loading = true;
			const { user } = await register({ email, name, password }, { autoLogin: true });
			onSuccess(user);
		} catch (error) {
			loading = false;
			const problem = toProblemDetails(error);

			if (isProblemType(problem, 'auth/email-taken')) {
				const hint = $t('errors.auth.email-taken', { values: { email } });
				(
					document.querySelector('input[name="email"]') as HTMLInputElement | null
				)?.setCustomValidity(hint);
				return;
			}
			errorMessage = getTranslatedError($t, error, { values: { email } });
		}
	};

	const onsubmit = async (event: Event) => {
		event.preventDefault();

		if (password !== confirmPassword) {
			const confirmPasswordInput = document.querySelector(
				'input[name="confirm-password"]'
			) as HTMLInputElement;
			confirmPasswordInput.setCustomValidity($t('passwords_must_match'));
			return;
		}

		await handleRegister();
	};
</script>

<AuthPageLayout title={$t('sign_up_please')}>
	<form {onsubmit} class="mt-6 space-y-4">
		<Alert type="error" message={errorMessage} />

		<!-- Email -->
		<div>
			<label for="email" class="label">
				<span class="label-text">{$t('email_address')}</span>
			</label>
			<input
				name="email"
				type="email"
				placeholder={$t('email_placeholder')}
				bind:value={email}
				oninput={(self) => {
					(self.target as HTMLInputElement).setCustomValidity('');
				}}
				class="input-bordered validator input w-full"
				required
			/>
			<p class="validator-hint hidden">
				{$t('errors.auth.email-taken', { values: { email } })}
			</p>
		</div>

		<!-- Username -->
		<div>
			<label for="name" class="label">
				<span class="label-text">{$t('name')}</span>
			</label>
			<input
				name="name"
				type="text"
				placeholder={$t('name_placeholder')}
				bind:value={name}
				class="input-bordered input w-full"
				required
			/>
		</div>

		<!-- Password -->
		<div>
			<label for="password" class="label">
				<span class="label-text">{$t('password')}</span>
			</label>
			<input
				name="password"
				type="password"
				placeholder="••••••••"
				bind:value={password}
				class="input-bordered validator input w-full"
				required
				minlength={passwordMinLength}
				title="Must be more than {passwordMinLength} characters"
			/>
			<p class="validator-hint hidden">
				{$t('password_min_length_hint', { values: { count: passwordMinLength } })}
			</p>
		</div>

		<!-- Confirm password -->
		<div>
			<label for="confirm-password" class="label">
				<span class="label-text">{$t('confirm_password')}</span>
			</label>
			<input
				name="confirm-password"
				type="password"
				placeholder="••••••••"
				bind:value={confirmPassword}
				class="input-bordered validator input w-full"
				required
			/>
			<p class="validator-hint hidden">
				{$t('passwords_must_match')}
			</p>
		</div>

		<!-- Register button -->
		<Button {loading} loadingText={$t('sign_up_loading')} type="submit" class="w-full btn-primary">
			{$t('sign_up')}
		</Button>

		<div class="divider">{$t('or_continue_with')}</div>

		<!-- Social logins -->
		<div class="grid grid-cols-2 gap-3">
			<button type="button" class="btn w-full gap-2 btn-outline">
				<img
					src="https://www.svgrepo.com/show/303108/google-icon-logo.svg"
					alt="Google"
					class="h-5 w-5"
				/>
				Google
			</button>
			<button type="button" class="btn w-full gap-2 btn-outline">
				<img src="https://www.svgrepo.com/show/217753/github.svg" alt="GitHub" class="h-5 w-5" />
				GitHub
			</button>
		</div>
	</form>

	<p class="mt-6 text-center text-sm">
		{$t('already_have_account')}
		<a href="/auth/login" class="link link-primary">{$t('sign_in')}</a>
	</p>
</AuthPageLayout>
