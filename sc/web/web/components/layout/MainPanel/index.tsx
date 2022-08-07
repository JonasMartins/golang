import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Group, Stack, Title, useMantineColorScheme } from "@mantine/core";
import React, { useCallback, useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/app";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import MessageComp from "@/components/layout/Messages";
import { MessageType } from "@/features/types/chat";
import CreateMessageForm from "@/components/form/CreateMessage";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.chat.value);
	const user = useSelector((state: RootState) => state.user.value);
	const [messages, setMessages] = useState<MessageType[]>([]);
	const { colorScheme } = useMantineColorScheme();

	const dark = colorScheme === "dark";
	const handleSetTitle = (): string => {
		if (user && chatFocused) {
			return GetChatTitle(chatFocused, user.id);
		}
		return "Unknown";
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
					spacing="lg"
					mt="lg"
					p={"lg"}
				>
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
				</Stack>
			</Stack>
			<CreateMessageForm />
		</Stack>
	);

	return content;
};

export default MainPanel;
