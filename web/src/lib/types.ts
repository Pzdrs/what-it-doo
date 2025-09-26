import type { UUID } from "crypto"

type User = {
    id: UUID;
    name: string;
    avatarUrl?: string;
}

export class Chat {
    id: number;
    title?: string;
    lastMessage?: ChatMessage;
    participants: User[];
    nicknames?: Record<UUID, string>;
    avatarUrl?: string;

    constructor(
        id: number,
        participants: User[],
        options: {
            title?: string;
            lastMessage?: ChatMessage;
            nicknames?: Record<UUID, string>;
            avatarUrl?: string;
        } = {}
    ) {
        this.id = id;
        this.participants = participants;
        this.title = options.title;
        this.lastMessage = options.lastMessage;
        this.nicknames = options.nicknames;
        this.avatarUrl = options.avatarUrl;
    }

    getTitle(): string {
        return this.title ?? 'Untitled chat';
    }

    getStatus(): boolean {
        return true;
    }

    getAvatarUrl(): string {
        return this.avatarUrl ?? 'https://randomuser.me/api/portraits/men/1.jpg';
    }
}

type ChatMessage = {
    id: UUID;
    chatId: number;
    content: string;
    sender: User;
    timestamp: Date;
    deliveredAt?: Date;
    readAt?: Date;
}

export type {
    User,
    ChatMessage
};