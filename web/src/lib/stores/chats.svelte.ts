import {
	getChatById,
	type DtoChat,
	type DtoChatMessage,
	type DtoUserDetails
} from '$lib/api/client';
import type { UUID } from 'crypto';

class MessagingStore {
	chats = $state<DtoChat[]>([]);
	messages = $state<Record<number, DtoChatMessage[]>>({}); // key = chatId, value = messages array

	_currentChatId = $state<number | null>(null);

	currentChat = $derived.by(() => {
		if (this._currentChatId === null) return null;
		return this.chats.find((chat) => chat.id === this._currentChatId) || null;
	});

	currentMessages = $derived.by(() => {
		if (this._currentChatId === null) return [];
		return this.messages[this._currentChatId] || [];
	});

	// List of all unique participants across all chats
	allParticipants = $derived.by(() => {
		const participantIds = new Set<UUID>();
		return this.chats
			.flatMap((chat) => chat.participants)
			.filter((p) => {
				if (participantIds.has(p.id)) return false;
				participantIds.add(p.id);
				return true;
			});
	});

	getChat(chatId: number): DtoChat | null {
		return this.chats.find((chat) => chat.id === chatId) || null;
	}

	getParticipant(user_id: UUID): DtoUserDetails | null {
		for (const chat of this.chats) {
			const participant = chat.participants.find((p) => p.id === user_id);
			if (participant) {
				return participant;
			}
		}
		return null;
	}

	async setCurrentChat(chatId: number) {
		this._currentChatId = chatId;
		if (this.chats.length === 0) {
			await getChatById(chatId).then((chat) => {
				this.chats.push(chat);
			});
		}
	}
}

export const messagingStore = new MessagingStore();
