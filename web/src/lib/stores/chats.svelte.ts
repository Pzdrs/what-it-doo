import { Chat, type ChatMessage } from "$lib/types";
import { getUser } from "./user.svelte";

let chats = $state<Chat[]>([
    new Chat(1, [
        {
            id: 'c175258b-88fe-4035-8c0b-c78c49ffee67',
            name: 'Alice',
            avatarUrl: 'https://randomuser.me/api/portraits/women/1.jpg'
        }, getUser()
    ], {
        title: 'Chat 1', lastMessage: {
            id: crypto.randomUUID(),
            chatId: 1,
            sender: getUser(),
            content: 'Hello, how are you?',
            timestamp: new Date(Date.now() - 300 * 1000), // 5 minutes ago
            readAt: new Date(Date.now() - 200 * 1000) // 3 minutes ago
        },
        typingUsers: []
    }),
    new Chat(2, [], {
        title: 'Chat 2',
        lastMessage: {
            id: crypto.randomUUID(),
            chatId: 2,
            sender: getUser(),
            content: "What's up?",
            timestamp: new Date(Date.now() - 3600 * 1000), // 1 hour ago
            readAt: new Date(Date.now()),
        },
    }),
    new Chat(3, [], {
        lastMessage: {
            id: crypto.randomUUID(),
            chatId: 3,
            sender: getUser(),
            content: "Let's catch up!",
            timestamp: new Date('2025-01-15T15:30:00Z')
        },
    })
]);
let messages = $state<ChatMessage[]>([
    {
        id: crypto.randomUUID(),
        chatId: 1,
        sender: {
            id: crypto.randomUUID(),
            name: 'Bob',
            avatarUrl: 'https://randomuser.me/api/portraits/men/1.jpg'
        },
        content: 'Hello, how are you? lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
        timestamp: new Date(Date.now() - 600 * 1000) // 10 minutes ago
    },
    {
        id: crypto.randomUUID(),
        chatId: 1,
        sender: {
            id: 'c175238b-88fe-4035-8c0b-c78c49ffee67',
            name: 'Alice',
            avatarUrl: 'https://randomuser.me/api/portraits/women/1.jpg'
        },
        content: 'Hi Bob! I am good, thanks!',
        //timestamp: new Date(Date.now() - 300 * 1000) // 5 minutes ago
    }
]);

export const getChats = () => {
    return chats;
}

export const getChat = (chatId: number): Chat | undefined => {
    return chats.find(chat => chat.id === chatId);
}

export const sendMessage = (message: ChatMessage) => {
    messages.push(message);
}

export const getMessages = (chatId: number) => {
    return messages.filter(message => message.chatId === chatId);
}