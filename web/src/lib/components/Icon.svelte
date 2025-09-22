<script lang="ts">
	import type { IconProps } from '../../types.js';
	import type { HTMLAttributes } from 'svelte/elements';

	const {
		size = '1em',
		viewBox = '0 0 24 24',
		class: className = '',
		strokeColor = 'transparent',
		strokeWidth = 2,
		role = 'img',
		title,
		icon,
		color = 'currentColor',
		description,
		...restProps
	}: IconProps & HTMLAttributes<EventTarget> = $props();
</script>

<svg
	width={size}
	height={size}
	{viewBox}
	class={className}
	stroke={strokeColor}
	stroke-width={strokeWidth}
	{role}
	{...restProps}
>
	{#if title}
		<title>{title}</title>
	{/if}
	{#if description}
		<desc>{description}</desc>
	{/if}
	<path d={typeof icon === 'string' ? icon : icon.path} fill={color} />
</svg>

<style>
	svg {
		transition: transform 0.2s ease;
	}
</style>
