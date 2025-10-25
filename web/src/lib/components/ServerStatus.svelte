<script lang="ts">
	import { getServerInfo, type DtoServerInfo } from '$lib/api/client';
	import { websocketStore } from '$lib/stores/websocket.svelte';
	import { t } from 'svelte-i18n';

	let info: DtoServerInfo = $state({});

	$effect(() => {
		getServerInfo()
			.then((res) => {
				info = res;
			})
			.catch((err) => {});

	});
</script>

<div class="flex justify-between px-2">
	<div class="flex items-center space-x-2">
		{#if websocketStore.connected}
			<div class="inline-grid *:[grid-area:1/1]">
				<div class="status status-success animate-ping"></div>
				<div class="status status-success"></div>
			</div>
			<span class="mr-5">
				{$t('server_status_online')}
			</span>
		{:else}
			<div class="inline-grid *:[grid-area:1/1]">
				<div class="status status-error animate-ping"></div>
				<div class="status status-error"></div>
			</div>
			<span class="mr-5">
				{$t('server_status_offline')}
			</span>
		{/if}
	</div>
	<span>{info.version}</span>
</div>
