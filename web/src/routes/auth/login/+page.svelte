<script>
	import AuthPageLayout from '$lib/components/layout/AuthPageLayout.svelte';
	import { t } from 'svelte-i18n';

	let email = '';
	let password = '';
	let remember = false;
</script>

<AuthPageLayout>
	<div class="flex min-h-screen items-center justify-center bg-base-200">
		<div class="w-full max-w-xl">
			<div class="card bg-base-100 shadow-xl">
				<div class="card-body p-10">
					<!-- Logo -->
					<div class="flex flex-col items-center">
						<span class="mb-5 text-3xl font-bold text-primary">what it doo</span>
						<h2 class="mt-4 text-2xl font-bold">{$t("sign_in_please")}</h2>
					</div>

					<!-- Form -->
					<form class="mt-6 space-y-4">
						<div>
							<label class="label">
								<span class="label-text">{$t("email_address")}</span>
							</label>
							<input
								type="email"
								placeholder="{$t("email_placeholder")}"
								bind:value={email}
								class="input-bordered input w-full"
								required
							/>
						</div>

						<div>
							<label class="label">
								<span class="label-text">{$t("password")}</span>
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
								<input type="checkbox" bind:checked={remember} class="checkbox checkbox-primary" />
								<span class="label-text">{$t("remember_me")}</span>
							</label>
							<a href="#" class="link text-sm link-primary">{$t("forgot_password")}</a>
						</div>

						<button type="submit" class="btn w-full btn-primary">{$t("sign_in")}</button>

						<div class="divider">{$t("or_continue_with")}</div>

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
								<img
									src="https://www.svgrepo.com/show/217753/github.svg"
									alt="GitHub"
									class="h-5 w-5"
								/>
								GitHub
							</button>
						</div>
					</form>

					<p class="mt-6 text-center text-sm">
						{$t("no_account")}
						<a href="/auth/register" class="link link-primary">{$t("sign_up")}</a>
					</p>
				</div>
			</div>
		</div>
	</div>
</AuthPageLayout>
