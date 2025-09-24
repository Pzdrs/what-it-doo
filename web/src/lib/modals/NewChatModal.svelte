<script>
	import TagsInput from '$lib/components/TagsInput.svelte';
</script>

<dialog id="new-chat-dialog" class="modal">
	<div class="modal-box w-6/12 max-w-none">
		<h3 class="text-lg font-bold">Start a new chat</h3>
		<p class="py-4">Press ESC key or click outside to close</p>
		<TagsInput />
		<div class="modal-action">
			<a href="#" class="btn">Yay!</a>
		</div>
	</div>

	<form method="dialog" class="modal-backdrop">
		<button>close</button>
	</form>
</dialog>
