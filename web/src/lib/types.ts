import type { UUID } from "crypto"

type User = {
    id: UUID;
    name: string;
    avatarUrl?: string;
}

type Chat = {
    id: number;
    title: string;
    lastMessage: ChatMessage;
    participants: User[];
}

type ChatMessage = {
    id: UUID;
    chatId: number;
    content: string;
    timestamp: Date;
    deliveredAt?: Date;
    readAt?: Date;
}

export type { Chat, ChatMessage };