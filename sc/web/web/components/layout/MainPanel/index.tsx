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
} from "@mantine/core";
import React, { useCallback, useEffect, useState, useRef } from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/app";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import MessageComp from "@/components/layout/Messages";
import { MessageType } from "@/features/types/chat";
import CreateMessageForm from "@/components/form/CreateMessage";
import { ChevronsDown } from "tabler-icons-react";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
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
		return "Unknown";
	};

	const scrollToBottom = () => {
		if (viewport !== undefined) {
			viewport.current?.scrollTo({ top: viewport.current.scrollHeight, behavior: "smooth" });
		} else {
			console.log("undefined ? ");
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
		if (!chatFocused || !messageHasBeenAdded.messageCount) return;

		if (chatFocused?.base.id === messageHasBeenAdded.chatId) {
			console.log("has added");
			if (viewport !== undefined) {
				viewport.current?.scrollTo({
					top: viewport.current.scrollHeight,
					behavior: "smooth",
				});
			}
		}
	}, [messageHasBeenAdded.messageCount]);

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
					p={"lg"}
					m="lg"
				>
					<ScrollArea style={{ height: "68vh" }} viewportRef={viewport}>
						{messages.map((x, i) => (
							<>
								<MessageComp
									key={i}
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
				<ActionIcon onClick={scrollToBottom}>
					<ChevronsDown />
				</ActionIcon>
				<Grid align="center" gutter="xs" grow>
					{/* <Grid.Col offset={1} span={1}>
					<Group>
						<ActionIcon>
							<MoodSmile />
						</ActionIcon>
					</Group>
				</Grid.Col> */}
					<Grid.Col span={12}>
						<CreateMessageForm />
					</Grid.Col>
				</Grid>
			</Stack>
		</Stack>
	);

	return content;
};

export default MainPanel;
