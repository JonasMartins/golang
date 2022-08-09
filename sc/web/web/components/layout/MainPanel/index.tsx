import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import {
	Group,
	Stack,
	Title,
	useMantineColorScheme,
	ScrollArea,
	ActionIcon,
	Grid,
	Tooltip,
} from "@mantine/core";
import React, { useCallback, useEffect, useState, useRef } from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/app";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import MessageComp from "@/components/layout/Messages";
import { MessageType } from "@/features/types/chat";
import CreateMessageForm from "@/components/form/CreateMessage";
import { ChevronsDown, ChevronsUp, Typography, TypographyOff } from "tabler-icons-react";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
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
		return "Unknown";
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
			const length = chatFocused.Messages.length - 1;
			for (let i = 0; i <= length; i++) {
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
					<Title order={4}>{handleSetTitle()}</Title>
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
					<ScrollArea style={{ height: "57vh" }} viewportRef={viewport}>
						{messages.map((x, i) => (
							<>
								<MessageComp
									key={x.base.id}
									message={x}
									nextMessageDate={
										i + 1 < messages.length
											? new Date(messages[i + 1].base.createdAt)
											: new Date(x.base.createdAt)
									}
								/>
							</>
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
							size="lg"
							variant="outline"
						>
							<ChevronsDown />
						</ActionIcon>
					</Tooltip>

					<Tooltip withArrow label="Scroll to top">
						<ActionIcon
							ml={"sm"}
							onClick={scrollToTop}
							radius="lg"
							size="lg"
							variant="outline"
						>
							<ChevronsUp />
						</ActionIcon>
					</Tooltip>
					<Tooltip withArrow label="Enter new lines messages">
						<ActionIcon
							ml={"sm"}
							onClick={() => {
								setNewLinesMessage(!newLinesMessage);
							}}
							radius="lg"
							size="lg"
							variant="outline"
						>
							{newLinesMessage ? <Typography /> : <TypographyOff />}
						</ActionIcon>
					</Tooltip>
				</Group>
				<Grid align="center" gutter="xs" grow>
					<Grid.Col span={12}>
						<CreateMessageForm showSubmitButton={newLinesMessage} />
					</Grid.Col>
				</Grid>
			</Stack>
		</Stack>
	);

	return content;
};

export default MainPanel;
