<script lang="ts">
    type ChatMessageOrigin = 'us' | 'them';

    interface Props {
        message: string;
        origin: ChatMessageOrigin;
    }

    let {message, origin = 'them'}: Props = $props();
</script>

<div class="my-2 flex" class:justify-end={origin === 'us'} class:justify-start={origin === 'them'}>
    <div class="max-w-xs rounded-lg p-3 text-white" class:bg-blue-500={origin === 'us'} class:bg-gray-500={origin === 'them'}>
        {message}
    </div>
</div>