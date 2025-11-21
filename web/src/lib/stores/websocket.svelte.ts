import type { DtoChatMessage } from '$lib/api/client';
import type { UUID } from 'crypto';
import { messagingStore } from './chats.svelte';

class WebsocketStore {
	// Used for things like showing online/offline server status
	connected = $state(false);
}

export const websocketStore = new WebsocketStore();

let websocket: WebSocket | null = null;
let reconnectTimeout = 1000;

const onMessage = (event: MessageEvent) => {
	const message = JSON.parse(event.data);
	switch (message.type) {
		case 'message_ack': {
			const { message_id, temp_id, sent_at } = message.data as {
				message_id: number;
				temp_id: number;
				sent_at: string;
			};
			messagingStore.acknowledgeMessage(temp_id, message_id, new Date(sent_at));
			break;
		}
		case 'new_message': {
			const { chat_id, message: newMessage } = message.data as {
				chat_id: number;
				message: DtoChatMessage;
			};
			messagingStore.addIncomingMessage(newMessage, chat_id);
			break;
		}
		case 'typing': {
			const { chat_id, user_id, typing } = message.data as {
				chat_id: number;
				user_id: UUID;
				typing: boolean;
			};
			if (typing) {
				messagingStore.addTypingUser(chat_id, user_id);
			} else {
				messagingStore.removeTypingUser(chat_id, user_id);
			}
			break;
		}
		case 'presence_change': {
			const { user_id, online } = message.data as {
				user_id: UUID;
				online: boolean;
			};
			messagingStore.updateUserPresence(user_id, online);
			break;
		}

		default:
			console.warn('Unknown message type:', message.type);
	}
};

export const openWebSocketConnection = () => {
	if (websocket && websocket.readyState === WebSocket.OPEN) {
		return;
	}

	try {
		websocket = new WebSocket('/api/v1/ws');

		websocket.onopen = () => {
			websocketStore.connected = true;
		};

		websocket.onclose = () => {
			websocketStore.connected = false;

			setTimeout(openWebSocketConnection, reconnectTimeout);
			reconnectTimeout *= 2; // Exponential backoff
		};

		websocket.onerror = (err) => {
			console.error('WebSocket connection error:', err);
			websocketStore.connected = false;
		};

		websocket.onmessage = onMessage;
	} catch (error) {
		console.error('Error opening WebSocket connection:', error);
	}
};

export const closeWebSocketConnection = () => {
	if (websocket) {
		websocket.close();
		websocket = null;
	}
};

export const sendWebSocketMessage = (type: string, data: unknown) => {
	websocket?.send(
		JSON.stringify({
			type,
			data
		})
	);
};
