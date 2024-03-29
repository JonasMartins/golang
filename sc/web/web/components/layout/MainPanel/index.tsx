import { RootState } from "@/app";
import CreateMessageForm from "@/components/form/CreateMessage";
import MessageComp from "@/components/layout/Messages";
import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import SettingsModal from "@/components/modal/Settings";
import GeneralMutationsAlert from "@/components/notifications/alert/Alert";
import { MessageType } from "@/features/types/chat";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import {
	ActionIcon,
	Avatar,
	Grid,
	Group,
	Indicator,
	ScrollArea,
	Stack,
	Title,
	Tooltip,
	useMantineColorScheme,
	Box,
} from "@mantine/core";
import { IconChevronsDown } from "@tabler/icons";
import React, { useCallback, useEffect, useRef, useState } from "react";
import { useSelector } from "react-redux";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const alertOpen = useSelector((state: RootState) => state.persistedReducer.alert.open);
	const messageHasBeenAdded = useSelector(
		(state: RootState) => state.persistedReducer.chat.hasAddedMessage
	);

	const user = useSelector((state: RootState) => state.persistedReducer.user.value);
	const [messages, setMessages] = useState<MessageType[]>([]);
	const { colorScheme } = useMantineColorScheme();

	const viewport = useRef<HTMLDivElement>(null);

	const dark = colorScheme === "dark";
	const handleSetTitle = (): string => {
		if (user && chatFocused) {
			return GetChatTitle(chatFocused, user.id);
		}
		return "";
	};

	const scrollToBottom = () => {
		if (viewport !== undefined) {
			viewport.current?.scrollTo({
				top: viewport.current.scrollHeight,
				behavior: "smooth",
			});
		}
	};

	const handleSettingMessagesToState = useCallback(() => {
		if (chatFocused) {
			for (let i = 0; i < chatFocused.Messages.length; i++) {
				const m: MessageType = chatFocused.Messages[i];
				setMessages(x => [...x, m]);
			}
		}
	}, [chatFocused]);

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
		if (messageHasBeenAdded) {
			scrollToBottom();
		}
	}, [messages.length, messageHasBeenAdded]);

	useEffect(() => {
		return () => {
			setMessages([]);
		};
	}, []);

	const content = (
		<Stack mb="sm" justify="space-between">
			<Stack>
				<Group
					sx={theme => ({
						backgroundColor: dark ? theme.colors.dark[5] : theme.colors.gray[3],
					})}
					position="apart"
					p="sm"
				>
					<Group>
						<Indicator color="green" position="bottom-end">
							<Avatar size={"md"} radius={"lg"} />
						</Indicator>
						<Title order={4}>{handleSetTitle()}</Title>
					</Group>
					<Group>
						<ToggleTheme />
						<SettingsMenu />
					</Group>
				</Group>

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
					<ScrollArea style={{ height: "70vh" }} viewportRef={viewport} offsetScrollbars>
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
					</ScrollArea>
				</Stack>
			</Stack>
			<Stack>
				<Group align="flex-start" sx={{ paddingLeft: "10px" }}>
					<Tooltip withArrow label="Scroll to bottom">
						<ActionIcon
							ml={"sm"}
							onClick={scrollToBottom}
							radius="lg"
							size="sm"
							variant="outline"
							disabled={chatFocused ? false : true}
						>
							<IconChevronsDown />
						</ActionIcon>
					</Tooltip>
				</Group>
				<Grid align="center" gutter="xs" grow>
					<Grid.Col span={12}>
						<CreateMessageForm />
					</Grid.Col>
				</Grid>
			</Stack>
			<SettingsModal />
		</Stack>
	);

	return content;
};

export default MainPanel;
