import {
	Paper,
	Stack,
	Group,
	Title,
	Text,
	Sx,
	Avatar,
	Indicator,
	useMantineColorScheme,
	Badge,
} from "@mantine/core";
import type { NextPage } from "next";
import { ChatType } from "@/features/types/chat";
import { formatRelative } from "date-fns";
import { setFocusedChat, clearMessageAddedOnChangeChat } from "@/features/chat/chatSlicer";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/app";
import { useCallback, useEffect, useState } from "react";

interface ChatsSideBarProps {
	chat: ChatType;
	title: string;
}

const ChatsSideBar: NextPage<ChatsSideBarProps> = ({ chat, title }) => {
	const dispatch = useDispatch();
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const loggedUser = useSelector((state: RootState) => state.persistedReducer.user.value);
	const [unseenMessagesCount, setUnseenMessagesCount] = useState(0);
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

	const handleGetUnSeenMessagesCount = useCallback(() => {
		if (loggedUser) {
			let count = 0;
			chat.Messages.map(x => {
				if (!x.Seen?.includes(loggedUser.id)) {
					count++;
				}
			});
			setUnseenMessagesCount(count);
		}
	}, [chat, loggedUser]);

	useEffect(() => {
		handleGetUnSeenMessagesCount();
		return () => {
			setUnseenMessagesCount(0);
		};
	}, []);

	return (
		<Paper
			sx={dark ? darkSx : lightSx}
			shadow="md"
			p="xs"
			radius="md"
			withBorder
			onClick={() => {
				dispatch(clearMessageAddedOnChangeChat());
				dispatch(setFocusedChat(chat.base.id));
			}}
		>
			<Stack>
				<Group grow align="center" position="apart">
					<Group>
						<Indicator color="green" position="bottom-end">
							<Avatar radius={"xl"} />
						</Indicator>
						<Title order={6}>{`${title}`}</Title>
					</Group>
					<Group sx={{ justifyContent: "flex-end" }}>
						<Text size="xs" align="right" weight={100}>
							{formatRelative(new Date(chat.base.updatedAt), new Date())}
						</Text>
						{unseenMessagesCount > 0 ? (
							<Badge size="sm" color={"red"} variant="filled">
								{unseenMessagesCount}
							</Badge>
						) : (
							<></>
						)}
					</Group>
				</Group>
				<Text size="xs">{chat.Messages[0].Body}</Text>
			</Stack>
		</Paper>
	);
};

export default ChatsSideBar;
