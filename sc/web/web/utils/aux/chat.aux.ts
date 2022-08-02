import { ChatType } from "@/features/types/chat";
export const GetChatTitle = (chat: ChatType, loggedUserId: string): string => {
	let title = "Unknown";

	chat.Members.map(x => {
		if (x.base.id != loggedUserId) {
			title = x.name;
		}
	});

	return title;
};
