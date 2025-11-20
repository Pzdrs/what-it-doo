import {
	getChatById,
	getChatMessages,
	type DtoChat,
	type DtoChatMessage,
	type DtoUserDetails
} from '$lib/api/client';
import { INIT_LOAD_MESSAGES_COUNT, LOAD_OLDER_MESSAGES_COUNT } from '$lib/constants';
import type { UUID } from 'crypto';
import { SvelteSet } from 'svelte/reactivity';

class MessagingStore {
	chats = $state<DtoChat[]>([]);
	messages = $state<Record<number, DtoChatMessage[]>>({}); // key = chatId, value = messages array
	typingUsers = $state<Record<number, SvelteSet<UUID>>>({}); // key = chatId, value = set of user IDs typing

	_currentChatId = $state<number | null>(null);
	_fullyPopulatedChats = $state<number[]>([]);

	currentChat = $derived.by(() => {
		if (this._currentChatId === null) return null;
		return this.chats.find((chat) => chat.id === this._currentChatId) || null;
	});

	currentMessages = $derived.by(() => {
		if (this._currentChatId === null) return [];

		const msgs = this.messages[this._currentChatId] || [];

		// return a sorted copy
		return [...msgs].sort(
			(a, b) => new Date(a?.sent_at).getTime() - new Date(b?.sent_at).getTime()
		);
	});

	// List of all unique participants across all chats
	allParticipants = $derived.by(() => {
		const participantIds = new SvelteSet<UUID>();
		return this.chats
			.flatMap((chat) => chat.participants)
			.filter((p) => {
				if (participantIds.has(p?.id)) return false;
				participantIds.add(p.id);
				return true;
			});
	});

	currentTypingUsers = $derived.by(() => {
		if (this._currentChatId === null) return [];

		const typingIds = this.typingUsers[this._currentChatId];
		if (!typingIds) return [];

		// Build a map for fast lookup: UUID → User
		const participantMap = new Map(this.allParticipants.map((p) => [p.id, p]));

		// Convert typing IDs into user objects
		return [...typingIds]
			.map((id) => participantMap.get(id))
			.filter((u): u is DtoUserDetails => u !== undefined);
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

	async initLoadMessages(chatId: number) {
		if (this.messages[chatId]) return; // already loaded
		await getChatMessages(chatId, { limit: INIT_LOAD_MESSAGES_COUNT }).then((dto) => {
			this.messages[chatId] = dto.messages;

			if (!dto.has_more) {
				this._fullyPopulatedChats.push(chatId);
			}
		});
	}

	async loadOlderMessages() {
		let loadedSome = false;

		if (this._currentChatId === null || this._fullyPopulatedChats.includes(this._currentChatId))
			return false;

		const chatId = this._currentChatId;
		const currentMsgs = this.messages[chatId] || [];
		const oldestMsg = currentMsgs.reduce(
			(oldest, msg) => {
				return !oldest || new Date(msg.sent_at) < new Date(oldest.sent_at) ? msg : oldest;
			},
			null as DtoChatMessage | null
		);

		const before = oldestMsg ? new Date(oldestMsg.sent_at).toISOString() : undefined;
		await getChatMessages(chatId, {
			before,
			limit: LOAD_OLDER_MESSAGES_COUNT
		}).then((dto) => {
			loadedSome = dto.messages?.length > 0;
			const existingMsgs = this.messages[chatId] || [];
			// Prepend older messages
			this.messages[chatId] = [...dto.messages, ...existingMsgs];

			if (!dto.has_more) {
				this._fullyPopulatedChats.push(chatId);
			}
		});

		return loadedSome;
	}

	async setCurrentChat(chatId: number) {
		this._currentChatId = chatId;
		if (this.chats.length === 0) {
			await getChatById(chatId).then((chat) => {
				this.chats.push(chat);
			});
		}
	}

	sendMessageOptimistic(messageId: number, chatId: number, senderId: UUID, content: string) {
		this.messages[chatId].push({
			id: messageId,
			sender_id: senderId,
			content: content
		});
	}

	addIncomingMessage(message: DtoChatMessage, chatId: number) {
		if (!this.messages[chatId]) {
			this.messages[chatId] = [];
		}
		// Avoid duplicates
		if (!this.messages[chatId].some((msg) => msg.id === message.id)) {
			this.messages[chatId].push(message);
		}
	}

	acknowledgeMessage(tempId: number, messageId: number, sentAt: Date) {
		for (const chatId in this.messages) {
			const msgs = this.messages[chatId];
			const msgIndex = msgs.findIndex((msg) => msg.id === tempId);
			if (msgIndex !== -1) {
				msgs[msgIndex].id = messageId;
				msgs[msgIndex].sent_at = sentAt.toISOString();
				break;
			}
		}
	}

	addTypingUser(chatId: number, userId: UUID) {
		if (!this.typingUsers[chatId]) {
			this.typingUsers[chatId] = new SvelteSet<UUID>();
		}
		this.typingUsers[chatId].add(userId);
	}

	removeTypingUser(chatId: number, userId: UUID) {
		if (this.typingUsers[chatId]) {
			this.typingUsers[chatId].delete(userId);
		}
	}

	updateUserPresence(user_id: string, online: boolean) {
		for (const chat in this.chats) {
			for (const participant of this.chats[chat].participants) {
				if (participant.id === user_id) {
					participant.online = online;
				}
			}
		}
	}
}

export const messagingStore = new MessagingStore();
