export type ChatType = {
	__typename?: "Chat";
	base: {
		__typename?: "Base";
		id: string;
		updatedAt: any;
	};
	Members: Array<{
		__typename?: "User";
		name: string;
		base: {
			__typename?: "Base";
			id: string;
		};
	}>;
	Messages: Array<{
		__typename?: "Message";
		Body: string;
		ChatId: string;
		Seen?: Array<string> | null;
		base: {
			__typename?: "Base";
			id: string;
			createdAt: any;
		};
		Author: {
			__typename?: "User";
			name: string;
		};
	}>;
};

export type Title = {
	title: string;
};

export type ChatTypeTitled = ChatType | Title;

export type AuthorMessage = {
	__typename?: "User";
	name: string;
};

export type MessageType = {
	__typename?: "Message";
	Body: string;
	ChatId: string;
	Seen?: Array<string> | null;
	base: {
		__typename?: "Base";
		id: string;
		createdAt: any;
	};
	Author: {
		__typename?: "User";
		name: string;
	};
};

export type MessageSubscription = {
	__typename?: "Message";
	Body: string;
	ChatId: string;
	AuthorId: string;
	Seen?: Array<string> | null;
	base: {
		__typename?: "Base";
		id: string;
		createdAt: any;
	};
};
