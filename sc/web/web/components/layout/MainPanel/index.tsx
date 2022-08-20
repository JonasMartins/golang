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
import { IconChevronsDown, IconChevronsUp } from "@tabler/icons";
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
	const [newLinesMessage, setNewLinesMessage] = useState(false);

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

	const scrollToTop = () => {
		if (viewport !== undefined) {
			viewport.current?.scrollTo({
				top: 0,
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

	useEffect(() => {
		handleSettingMessagesToState();
		return () => {
			setMessages([]);
		};
	}, [chatFocused, handleSettingMessagesToState]);

	useEffect(() => {
		if (!chatFocused || !messageHasBeenAdded) return;

		if (chatFocused?.base.id === messageHasBeenAdded.ChatId) {
			setMessages(x => [...x, messageHasBeenAdded]);
		}
	}, [messageHasBeenAdded, chatFocused]);

	useEffect(() => {
		scrollToBottom();
	}, [messages]);

	useEffect(() => {
		return () => {
			setMessages([]);
		};
	}, []);

	const content = (
		<Stack mb="sm" justify="space-between" sx={{ height: "100vh" }}>
			<Stack>
				<Group
					sx={theme => ({
						backgroundColor: dark ? theme.colors.dark[5] : theme.colors.gray[3],
					})}
					position="apart"
					p="sm"
				>
					<Group>
						<Indicator color="green">
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
					<ScrollArea style={{ height: "70vh" }} viewportRef={viewport}>
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

					<Tooltip withArrow label="Scroll to top">
						<ActionIcon
							ml={"sm"}
							onClick={scrollToTop}
							radius="lg"
							size="sm"
							variant="outline"
							disabled={chatFocused ? false : true}
						>
							<IconChevronsUp />
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
