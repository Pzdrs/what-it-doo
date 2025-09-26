<script>
	import TagsInput from '$lib/components/TagsInput.svelte';
</script>

<dialog id="new-chat-dialog" class="modal">
	<div class="modal-box h-6/12 flex w-6/12 max-w-none flex-col">
		<div class="flex-grow">
			<h3 class="text-lg font-bold">Start a new chat</h3>
			<p class="py-4">Please select the people you want to chat with:</p>
			<TagsInput />
		</div>
		<div class="modal-action mt-max">
			<a href="#" class="btn">Start Chat</a>
		</div>
	</div>

	<form method="dialog" class="modal-backdrop">
		<button>close</button>
	</form>
</dialog>
