import { RootState } from "@/app";
import MessageComp from "@/components/layout/Messages";
import SettingsModal from "@/components/modal/Settings";
import GeneralMutationsAlert from "@/components/notifications/alert/Alert";
import { MessageType } from "@/features/types/chat";
import { Box, Stack, ScrollArea, Group, ActionIcon, Text, Indicator } from "@mantine/core";
import { IconChevronsDown } from "@tabler/icons";
import React, { useCallback, useEffect, useRef, useState } from "react";
import { useSelector } from "react-redux";

interface ChatProps {}

const Chat: React.FC<ChatProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const alertOpen = useSelector((state: RootState) => state.persistedReducer.alert.open);
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
	const [messages, setMessages] = useState<MessageType[]>([]);

	const handleSettingMessagesToState = useCallback(() => {
		if (chatFocused) {
			for (let i = 0; i < chatFocused.Messages.length; i++) {
				const m: MessageType = chatFocused.Messages[i];
				setMessages(x => [...x, m]);
			}
		}
	}, [chatFocused]);

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

	useEffect(() => {
		window.addEventListener("resize", setDimension);
		return () => {
			window.removeEventListener("resize", setDimension);
		};
	}, [sizeScreen]);

	useEffect(() => {
		return () => {
			setMessages([]);
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
				>
					<Group>
						<ActionIcon
							sx={{ position: "fixed", bottom: 120 }}
							onClick={scrollToBottom}
							radius="xl"
							size="xl"
							variant="outline"
							disabled={chatFocused ? false : true}
						>
							<Indicator
								offset={-7}
								position="top-start"
								label={5}
								color="#007c00"
								size={30}
							>
								<IconChevronsDown />
							</Indicator>
						</ActionIcon>

						<Stack sx={{ flexGrow: 1, marginLeft: 50 }}>
							{messages.map((x, i) => (
								<MessageComp
									key={x.base.id}
									message={x}
									nextMessageDate={
										i + 1 < messages.length
											? new Date(messages[i + 1].base.createdAt)
											: new Date(x.base.createdAt)
									}
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
