import { RootState } from "@/app";
import MessageComp from "@/components/layout/Messages";
import SettingsModal from "@/components/modal/Settings";
import GeneralMutationsAlert from "@/components/notifications/alert/Alert";
import { MessageType } from "@/features/types/chat";
import { ActionIcon, Box, Group, Indicator, ScrollArea, Stack } from "@mantine/core";
import { IconChevronsDown } from "@tabler/icons";
import React, { useCallback, useEffect, useRef, useState } from "react";
import { useSelector } from "react-redux";

interface ChatProps {}

const Chat: React.FC<ChatProps> = () => {
	const alertOpen = useSelector((state: RootState) => state.persistedReducer.alert.open);
	const loggedUser = useSelector((state: RootState) => state.persistedReducer.user.value);
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);

	const [messages, setMessages] = useState<MessageType[]>([]);
	const [unseenMessagesCount, setUnseenMessagesCount] = useState(0);
	const [chatMembersIds, setChatMembersIds] = useState<string[]>([]);
	const [scrollPosition, onScrollPositionChange] = useState({ x: 0, y: 0 });
	const [hasUnSeenMessagesIndicator, setHasUnSeenMessagesIndicator] = useState("");
	const messageHasBeenAdded = useSelector(
		(state: RootState) => state.persistedReducer.chat.hasAddedMessage
	);

	const [sizeScreen, getSizeScreen] = useState({
		dynamicWidth: window.innerWidth,
		dynamicHeight: window.innerHeight,
	});

	const setDimension = () => {
		getSizeScreen({
			dynamicHeight: window.innerHeight,
			dynamicWidth: window.innerWidth,
		});
	};

	const viewport = useRef<HTMLDivElement>(null);

	/**
	 * 	Initial setup to set the focused chat messages
	 * 	to this component
	 */
	const handleSettingMessagesToState = useCallback(() => {
		let lastUnseenMessageId = true;
		if (chatFocused && loggedUser) {
			chatFocused.Members.map(x => {
				setChatMembersIds(prev => [...prev, x.base.id]);
			});

			let count = 0;
			for (let i = 0; i < chatFocused.Messages.length; i++) {
				const m: MessageType = chatFocused.Messages[i];
				if (!m.Seen?.includes(loggedUser.id)) {
					if (lastUnseenMessageId) {
						lastUnseenMessageId = false;
						setHasUnSeenMessagesIndicator(m.base.id);
					}
					count++;
				}
				setMessages(x => [...x, m]);
			}
			setUnseenMessagesCount(count);
		}
	}, [chatFocused, loggedUser]);

	/**
	 * Method that scrolls the scroll area to absolute bottom
	 */
	const scrollToBottom = () => {
		if (viewport !== undefined) {
			viewport.current?.scrollTo({
				top: viewport.current.scrollHeight,
				behavior: "smooth",
			});
		}
	};
	/**
	 * Handles initial setup, this is triggered when the user
	 * switch chats
	 */
	useEffect(() => {
		handleSettingMessagesToState();
		return () => {
			setMessages([]);
			setUnseenMessagesCount(0);
			setHasUnSeenMessagesIndicator("");
			setChatMembersIds([]);
		};
	}, [chatFocused, handleSettingMessagesToState]);

	/**
	 *  Handle when a new message appears on messageHasBeenAdded
	 *  and updates the component state messages if the message
	 *  belongs to this chat
	 */
	useEffect(() => {
		if (!chatFocused || !messageHasBeenAdded) return;

		if (chatFocused?.base.id === messageHasBeenAdded.ChatId) {
			setMessages(x => [...x, messageHasBeenAdded]);
		}
	}, [messageHasBeenAdded, chatFocused]);

	/**
	 * UseEffect to monitoring the window's resizing
	 */
	useEffect(() => {
		window.addEventListener("resize", setDimension);
		return () => {
			window.removeEventListener("resize", setDimension);
		};
	}, [sizeScreen]);

	/**
	 * Cleaning states when components is about to end
	 */
	useEffect(() => {
		return () => {
			setMessages([]);
			setUnseenMessagesCount(0);
			setHasUnSeenMessagesIndicator("");
			setChatMembersIds([]);
		};
	}, []);

	const content = (
		<Stack mb={0} justify="space-between">
			<Stack
				sx={() => ({
					flexDirection: "column",
					justifyContent: "center",
				})}
			>
				{alertOpen && (
					<Box p="sm" m="sm">
						<GeneralMutationsAlert />
					</Box>
				)}
				<ScrollArea
					style={{ height: sizeScreen.dynamicHeight - 202 }}
					type="hover"
					viewportRef={viewport}
					offsetScrollbars
					onScrollPositionChange={onScrollPositionChange}
				>
					<Group>
						<ActionIcon
							sx={{ position: "fixed", bottom: 120 }}
							onClick={scrollToBottom}
							radius="lg"
							size="lg"
							variant="outline"
							disabled={chatFocused ? false : true}
						>
							<Indicator
								offset={-3}
								position="top-start"
								label={unseenMessagesCount}
								color="#007c00"
								size={30}
								hidden={chatFocused ? false : true}
							>
								<IconChevronsDown />
							</Indicator>
						</ActionIcon>

						<Stack sx={{ flexGrow: 1, marginLeft: 50 }}>
							{/**
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
							 */}
							{messages.map((x, i) => (
								<MessageComp
									key={x.base.id}
									message={x}
									nextMessageDate={
										i + 1 < messages.length
											? new Date(messages[i + 1].base.createdAt)
											: new Date(x.base.createdAt)
									}
									messagesUnSeenIndicator={
										hasUnSeenMessagesIndicator === x.base.id
									}
									chatMembersIds={chatMembersIds}
									scrollPosition={scrollPosition.y}
									currentPageHeight={sizeScreen.dynamicHeight - 202}
								/>
							))}
						</Stack>
					</Group>
				</ScrollArea>
			</Stack>
			<SettingsModal />
		</Stack>
	);

	return content;
};

export default Chat;
