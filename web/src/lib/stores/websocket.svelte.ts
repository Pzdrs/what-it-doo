class WebsocketStore {
	// Used for things like showing online/offline server status
	connected = $state(false);
}

export const websocketStore = new WebsocketStore();

let websocket: WebSocket | null = null;
let reconnectTimeout = 1000;

const onMessage = (event: MessageEvent) => {
    console.log('Received:', event.data);
    // Handle incoming messages
};

export const openWebSocketConnection = () => {
	try {
		websocket = new WebSocket('/api/v1/ws');

		websocket.onopen = () => {
			console.log('WebSocket connected');
			websocketStore.connected = true;
		};

		websocket.onclose = () => {
			console.log('WebSocket disconnected');
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

