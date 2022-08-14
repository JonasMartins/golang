import { Paper, Stack, Group, Title, Text, Sx, useMantineColorScheme } from "@mantine/core";
import type { NextPage } from "next";
import { ChatType } from "@/features/types/chat";
import { formatRelative } from "date-fns";
import { setFocusedChat, clearMessageAddedOnChangeChat } from "@/features/chat/chatSlicer";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/app";
import { useEffect } from "react";

interface ChatsSideBarProps {
	chat: ChatType;
	title: string;
}

const ChatsSideBar: NextPage<ChatsSideBarProps> = ({ chat, title }) => {
	const dispatch = useDispatch();
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const hasAddedMessage = useSelector(
		(state: RootState) => state.persistedReducer.chat.hasAddedMessage
	);
	const chatsFromReducer = useSelector((state: RootState) => state.persistedReducer.chat.chats);
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";

	let isChatInfocus = chatFocused?.base.id === chat.base.id;

	const lightSx: Sx = {
		cursor: "pointer",
		backgroundColor: isChatInfocus ? "#b2ffcb" : "#fff",
		"&:hover": {
			backgroundColor: isChatInfocus ? "#99e1ba" : "#dde0e2",
		},
	};

	const darkSx: Sx = {
		cursor: "pointer",
		backgroundColor: isChatInfocus ? "#07666c" : "#1A1B1E",
		"&:hover": {
			backgroundColor: isChatInfocus ? "#075257" : "#070707",
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
				dispatch(clearMessageAddedOnChangeChat());
				dispatch(setFocusedChat(chat.base.id));
			}}
		>
			<Stack>
				<Group grow align="center" position="apart">
					<Title order={6}>{`${title} (${chat.Messages.length})`}</Title>
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
