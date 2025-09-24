<script lang="ts">
	import type { Chat } from '$lib/types';
	import { format } from 'timeago.js';

	interface Props {
		chat: Chat;
		active?: boolean;
		userOnline?: boolean;
	}

	let { chat, active = false, userOnline = false }: Props = $props();
</script>

<li>
	<a href="/chat/{chat.id}" class:menu-active={active}>
		<div class="rounded-box flex items-center p-2 transition-colors duration-200">
			<div class="avatar" class:avatar-online={userOnline} class:avatar-offline={!userOnline}>
				<div class="w-12 rounded-full">
					<img
						src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p200x200_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=bW-Nk_SfhrEQ7kNvwEYj1ab&_nc_oc=AdmnK9O53ElfigOxXct-Vi8G0jm10Q64AR71Rb62wFxLKOt4gJJsq8UQPuRQm9IpmMI&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=iTU3jBr-XOTG0tPi5IzqBA&oh=00_AfZm3SJPNquCb3o8xQWx0bJc7fl10-3xZwVbUlw_WVim9g&oe=68D61548"
						alt="Chat Icon"
					/>
				</div>
			</div>
			<div class="ml-4">
				<h3 class="font-bold">{chat.title}</h3>
				<p class="text-sm">
					<span>{chat.lastMessage.content}</span> &middot;
					<span>{format(chat.lastMessage.timestamp)}</span>
				</p>
			</div>
		</div>
	</a>
</li>
