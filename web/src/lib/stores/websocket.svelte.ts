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
			const { temp_id, sent_at } = message.data as { temp_id: number; sent_at: string };
			messagingStore.acknowledgeMessage(temp_id, new Date(sent_at));
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
