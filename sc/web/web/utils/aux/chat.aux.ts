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

export const uuidv4Like = (): string => {
	return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, c => {
		var r = (Math.random() * 16) | 0,
			v = c == "x" ? r : (r & 0x3) | 0x8;
		return v.toString(16);
	});
};
