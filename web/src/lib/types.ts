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
    typingUsers?: UUID[];

    constructor(
        id: number,
        participants: User[],
        options: {
            title?: string;
            lastMessage?: ChatMessage;
            nicknames?: Record<UUID, string>;
            avatarUrl?: string;
            typingUsers?: UUID[];
        } = {}
    ) {
        this.id = id;
        this.participants = participants;
        this.title = options.title;
        this.lastMessage = options.lastMessage;
        this.nicknames = options.nicknames;
        this.avatarUrl = options.avatarUrl;
        this.typingUsers = options.typingUsers;
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

    getTypingUsers(): User[] {
        return this.typingUsers?.map(id => this.participants.find(user => user.id === id)).filter(Boolean) as User[] || [];
    }
}

type ChatMessage = {
    id: UUID;
    chatId: number;
    content: string;
    sender: User;
    timestamp?: Date;
    deliveredAt?: Date;
    readAt?: Date;
}

export type {
    User,
    ChatMessage
};