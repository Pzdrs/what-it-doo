<script lang="ts">
	import { createChat } from "$lib/api/client";
	import { userStore } from "$stores/user.svelte";
	import { t } from "svelte-i18n";

	let emails = $state(['']);

	function addEmail() {
		emails = [...emails, ''];
	}

	function removeEmail(index: number) {
		emails = emails.filter((_, i) => i !== index);
	}

	function updateEmail(index: number, value: string) {
		emails[index] = value;
	}

	async function startChat() {
		let validEmails = emails.filter((e) => e.trim() !== '');
		validEmails.push(userStore.user?.email);
		await createChat({ participants: validEmails });
		
		(document.getElementById('new-chat-dialog') as HTMLDialogElement)?.close();
	}
</script>

<dialog id="new-chat-dialog" class="modal" onclose={() => (emails = [''])}>
	<div class="modal-box h-6/12 w-6/12 max-w-none flex flex-col">
		<div class="flex-grow">
			<h3 class="text-lg font-bold">{$t('start_chat_long')}</h3>
			<p class="py-4">{$t('enter_emails_prompt')}</p>

			<div class="flex flex-col gap-2">
				{#each emails as email, i}
					<div class="flex gap-2 items-center">
						<input
							type="email"
							placeholder="{$t('email_placeholder')}"
							class="input input-bordered w-full"
							bind:value={emails[i]}
							oninput={(e) => updateEmail(i, e.target.value)}
							required
						/>
						{#if emails.length > 1}
							<button
								type="button"
								class="btn btn-error btn-sm"
								onclick={() => removeEmail(i)}
							>
								{$t('remove')}
							</button>
						{/if}
					</div>
				{/each}

				<button type="button" class="btn btn-outline mt-2" onclick={addEmail}>
					{$t('add_another_email')}
				</button>
			</div>
		</div>

		<div class="modal-action">
			<button class="btn btn-primary" type="button" onclick={startChat}>
				{$t('start_chat')}
			</button>
		</div>
	</div>

	<form method="dialog" class="modal-backdrop">
		<button>close</button>
	</form>
</dialog>
