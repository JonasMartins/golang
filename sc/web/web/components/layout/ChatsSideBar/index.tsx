import { Stack, Title } from "@mantine/core";
import type { NextPage } from "next";

type ChatType = {
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
		Seen?: Array<string> | null;
		base: {
			__typename?: "Base";
			createdAt: any;
		};
		Author: {
			__typename?: "User";
			name: string;
		};
	}>;
};

interface ChatsSideBarProps {
	chat: ChatType;
	currentUserId: string;
}

const ChatsSideBar: NextPage<ChatsSideBarProps> = ({ chat, currentUserId }) => {
	return (
		<Stack>
			<Title>{chat.base.id}</Title>
		</Stack>
	);
};

export default ChatsSideBar;
