/* eslint-disable react-hooks/exhaustive-deps */
import { RootState } from "@/app";
import { MessageType } from "@/features/types/chat";
import {
	ActionIcon,
	Box,
	Divider,
	Group,
	Paper,
	Stack,
	Sx,
	Text,
	useMantineColorScheme,
} from "@mantine/core";
import { IconCalendar, IconChevronDown, IconEye } from "@tabler/icons";
import { formatRelative } from "date-fns";
import type { NextPage } from "next";
import { useCallback, useEffect, useRef, useState } from "react";
import { useSelector } from "react-redux";

interface MessagesProps {
	message: MessageType;
	nextMessageDate: Date;
	messagesUnSeenIndicator: boolean;
	chatMembersIds: string[];
	scrollPosition: number;
	currentPageHeight: number;
}

const DAY_IN_MILI = 86400000;
/**
 *	Props:
 * 	message: the main message object to be rendered
 * 	nextMessageDate: The date from the next message, if there is no
 * 	next message them it comes the current message creation date
 * 	messageUnseenIndicator: Boolean, true if the current message is unread
 * 	by the current logged user
 * 	chatMembersIds: all the chat members ids
 *  scrollPosition: a number with the current y index position so it
 * 	can calculate if this message is visible during the scroll act
 * 	currentPageHeight: the page current height minus the navbat and footer
 * 	so it is the exacly main messages panel height
 *
 * @param param0
 * @returns
 */
const Messages: NextPage<MessagesProps> = ({
	message,
	nextMessageDate,
	messagesUnSeenIndicator,
	chatMembersIds,
	scrollPosition,
	currentPageHeight,
}) => {
	const refEl = useRef<HTMLDivElement>(null);
	const { colorScheme } = useMantineColorScheme();
	const [messageSeen, setMessageSeen] = useState(false);
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);

	const dark = colorScheme === "dark";
	let messageAuthor = message.Author.name === user?.name;

	/**
	 *  Initial function to determine if a message has been
	 * 	seen by every chat member
	 * @returns true if every member of the chat id is inside
	 * message's seen array
	 */
	const handleSeenMessageIndicator = (): boolean => {
		for (var i = 0; i < chatMembersIds.length; i++) {
			if (!message.Seen?.includes(chatMembersIds[i])) {
				return false;
			}
		}
		return true;
	};

	/**
	 *  Updating the state that indicates if this message has been seen,
	 *  if the message reference positin is visible in the screen, and this
	 *  state is false, meaning the user havent seen the message, then will
	 * 	be updated to true, making the eye icon turn blue
	 */
	const handleUpdateSeenMessageLocally = useCallback(() => {
		if (!messageSeen) {
			setMessageSeen(true);
		}
	}, [messageSeen]);

	useEffect(() => {
		if (!refEl) {
			return;
		}

		if (refEl.current) {
			/**
			 *  If the message screen position y index margin to top is inside
			 * 	panel visible position, then it will call a method to update current
			 *  seen state, indicating if every chat member has seen the message
			 */
			if (refEl.current && refEl.current.offsetTop < scrollPosition + currentPageHeight) {
				handleUpdateSeenMessageLocally();
			}
		}
	}, [scrollPosition, refEl]);

	useEffect(() => {
		setMessageSeen(handleSeenMessageIndicator());
	}, [chatMembersIds]);

	return (
		<Stack p="xs" mt="xs" ref={refEl} key-prop={message.base.id}>
			{messagesUnSeenIndicator && (
				<Paper radius="lg" p={1} withBorder>
					<Group position="center" align="center">
						<Text size={"sm"}>Last Unread Messages</Text>
						<ActionIcon>
							<IconChevronDown size={15} />
						</ActionIcon>
					</Group>
				</Paper>
			)}

			<Group sx={{ justifyContent: messageAuthor ? "flex-end" : "flex-start" }}>
				<Paper sx={getMessageStyle(dark, messageAuthor)} shadow="md" p="xs" radius={"lg"}>
					<Stack spacing="sm">
						<Text size="xs">{message.Body}</Text>
						<Group position="apart">
							<IconEye size={14} color={messageSeen ? "#05a3f4" : "#797c7b"} />
							<Text size="xs">
								{formatRelative(new Date(message.base.createdAt), new Date())}
							</Text>
						</Group>
					</Stack>
				</Paper>
			</Group>

			{nextMessageDate.getTime() - new Date(message.base.createdAt).getTime() >
				DAY_IN_MILI && (
				<Divider
					variant="dashed"
					size="sm"
					labelPosition="center"
					label={
						<>
							<IconCalendar size={20} />
							<Box ml={5}>{new Date(message.base.createdAt).toDateString()}</Box>
						</>
					}
				/>
			)}
		</Stack>
	);
};

export default Messages;

/**
 *
 * @param theme boolean, true if theme is dark, false otherwise
 * @param messageAuthor boolean, true if message author is the current logged user
 * @returns  the correct theme based on the current theme and message author
 */
const getMessageStyle = (theme: boolean, messageAuthor: boolean): Sx => {
	const lightSx: Sx = {
		cursor: "pointer",
		backgroundColor: messageAuthor ? "#b2ffcb" : "#fff",
		"&:hover": {
			backgroundColor: messageAuthor ? "#99e1ba" : "#dde0e2",
		},
	};

	const darkSx: Sx = {
		cursor: "pointer",
		backgroundColor: messageAuthor ? "#07666c" : "#1A1B1E",
		"&:hover": {
			backgroundColor: messageAuthor ? "#075257" : "#070707",
		},
	};

	return theme ? darkSx : lightSx;
};
