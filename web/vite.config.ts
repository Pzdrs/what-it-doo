import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const upstream = {
	target: process.env.WID_SERVER_URL || 'http://localhost:8080',
	changeOrigin: true,
	secure: true,
	ws: true
};

export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	server: {
		proxy: {
			'/api/v1': {
				...upstream,
				configure: (proxy, _options) => {
					proxy.on('proxyReqWs', (proxyReq, req, _socket, _options, _head) => {
						console.log('Proxying WebSocket Request to the Target:', req.url);
					});
					proxy.on('error', (err, _req, _res) => {
						console.log('proxy error', err);
					});
					proxy.on('proxyReq', (proxyReq, req, _res) => {
						console.log('Sending Request to the Target:', req.method, req.url);
					});
					proxy.on('proxyRes', (proxyRes, req, _res) => {
						console.log('Received Response from the Target:', proxyRes.statusCode, req.url);
					});
				}
			}
		}
	}
});
