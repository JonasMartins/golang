export type MessagesUnreadSet = {
	messageId: string;
	seen: string[];
};

export type CacheInputsSet = {
	userId: string;
	messageId: string;
};

export interface ChatMessagesUnreadSet {
	data: Record<string, Array<MessagesUnreadSet>>;
}

const chat1Messages: Array<MessagesUnreadSet> = [];
const chat2Messages: Array<MessagesUnreadSet> = [];

// will store recents inputs and avoid processing
// if the input was already sended
const recentUpdatedMessages: Array<CacheInputsSet> = [];

chat1Messages.push({
	messageId: "123",
	seen: ["11"],
});

chat1Messages.push({
	messageId: "124",
	seen: ["12"],
});

chat1Messages.push({
	messageId: "125",
	seen: ["11"],
});

chat1Messages.push({
	messageId: "126",
	seen: ["11"],
});

/////

chat2Messages.push({
	messageId: "223",
	seen: ["21"],
});
chat2Messages.push({
	messageId: "224",
	seen: ["22"],
});
chat2Messages.push({
	messageId: "225",
	seen: ["21"],
});
chat2Messages.push({
	messageId: "226",
	seen: ["22"],
});
const chats: ChatMessagesUnreadSet = {
	data: {
		"1": chat1Messages,
		"2": chat2Messages,
	},
};

const updateUnseenMessage = (chatId: string, messageId: string, userId: string) => {
	for (let j of recentUpdatedMessages) {
		if (j.messageId === messageId && j.userId === userId) {
			return;
		}
	}
	recentUpdatedMessages.push({ messageId, userId });

	for (let obj in chats.data) {
		if (obj === chatId) {
			// find the chat
			for (let i of chats.data[obj]) {
				// find the message
				if (i.messageId === messageId) {
					// add the user id who saw the message
					i.seen.push(userId);
				}
			}
		}
	}
};

export const run = () => {
	console.log("called");
	updateUnseenMessage("1", "123", "12");
	updateUnseenMessage("1", "123", "12");
	updateUnseenMessage("1", "123", "12");
	updateUnseenMessage("1", "124", "11");
	updateUnseenMessage("1", "124", "11");
	updateUnseenMessage("1", "125", "12");
	updateUnseenMessage("1", "125", "12");
	updateUnseenMessage("1", "126", "12");

	console.log(chat1Messages);
};
