import { Paper, Stack, Group, Title, Text, Sx, useMantineColorScheme } from "@mantine/core";
import type { NextPage } from "next";
import { ChatType } from "@/features/types/chat";
import { formatRelative } from "date-fns";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import { setFocusedChat } from "@/features/chat/chatSlicer";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/Redux";

interface ChatsSideBarProps {
	chat: ChatType;
	currentUserId: string;
}

const ChatsSideBar: NextPage<ChatsSideBarProps> = ({ chat, currentUserId }) => {
	const dispatch = useDispatch();
	const chatFocused = useSelector((state: RootState) => state.chat.value);
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";

	const lightSx: Sx = {
		cursor: "pointer",
		backgroundColor: chatFocused?.base.id === chat.base.id ? "#6ef0f8" : "#fff",
		"&:hover": {
			backgroundColor: chatFocused?.base.id === chat.base.id ? "#49dde6" : "#dde0e2",
		},
	};

	const darkSx: Sx = {
		cursor: "pointer",
		backgroundColor: chatFocused?.base.id === chat.base.id ? "#07858d" : "#1A1B1E",
		"&:hover": {
			backgroundColor: chatFocused?.base.id === chat.base.id ? "#07666c" : "#070707",
		},
	};

	return (
		<Paper
			sx={dark ? darkSx : lightSx}
			shadow="md"
			p="md"
			radius="md"
			withBorder
			onClick={() => {
				dispatch(setFocusedChat(chat));
			}}
		>
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
