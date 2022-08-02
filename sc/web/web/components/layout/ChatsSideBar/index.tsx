import { Paper, Stack, Group, Title, Text } from "@mantine/core";
import type { NextPage } from "next";
import { ChatType } from "@/features/types/chat";
import { formatRelative } from "date-fns";
import { GetChatTitle } from "@/utils/aux/chat.aux";

interface ChatsSideBarProps {
	chat: ChatType;
	currentUserId: string;
}

const ChatsSideBar: NextPage<ChatsSideBarProps> = ({ chat, currentUserId }) => {
	return (
		<Paper shadow="md" p="md" radius="md" withBorder>
			<Stack>
				<Group grow align="center" position="apart">
					<Title order={5}>{GetChatTitle(chat, currentUserId)}</Title>
					<Text size="xs" align="right" weight={100}>
						{formatRelative(new Date(chat.base.updatedAt), new Date())}
					</Text>
				</Group>
				<Text size="xs">{chat.Messages[0].Body}</Text>
			</Stack>
		</Paper>
	);
};

export default ChatsSideBar;
