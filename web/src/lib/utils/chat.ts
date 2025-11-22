import type { DtoChat, DtoUserDetails } from '$lib/api/client';

export const getTheOtherParticipant = (
	chat: DtoChat,
	myself: DtoUserDetails
): DtoUserDetails | null => {
	return chat.participants.find((p) => p.id !== myself.id) || null;
};

export const getOtherChatParticipants = (
	chat: DtoChat,
	myself: DtoUserDetails
): DtoUserDetails[] => {
	return chat.participants.filter((p) => p.id !== myself.id);
};

export const getGroupChatTitle = (
	chat: DtoChat,
	myself: DtoUserDetails,
	truncateAt: number
): string => {
	const participantNames = getOtherChatParticipants(chat, myself).map((p) => p.name.split(' ')[0]);

	const fullTitle = participantNames.join(', ');

	if (fullTitle.length <= truncateAt) {
		return fullTitle;
	}

	const sliceLength = Math.max(0, truncateAt - 3);

	return fullTitle.substring(0, sliceLength) + '...';
};
