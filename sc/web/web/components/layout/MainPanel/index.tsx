import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Group, Stack, Title, useMantineColorScheme } from "@mantine/core";
import React from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/app";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import MessageComp from "@/components/layout/Messages";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.chat.value);
	const user = useSelector((state: RootState) => state.user.value);
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";

	const handleSetTitle = (): string => {
		if (user && chatFocused) {
			return GetChatTitle(chatFocused, user.id);
		}
		return "Unknown";
	};

	const messagesListed = (
		<Stack
			sx={() => ({
				flexDirection: "column",
				justifyContent: "center",
			})}
			spacing="lg"
			mt="lg"
			p={"lg"}
		>
			{chatFocused?.Messages.slice(0)
				.reverse()
				.map(x => (
					<MessageComp key={x.base.createdAt} message={x} />
				))}
		</Stack>
	);

	const content = (
		<Stack
			sx={theme => ({
				backgroundColor: dark ? theme.colors.dark[5] : theme.colors.gray[2],
				height: "100vh",
			})}
		>
			<Group position="apart" m="md">
				<Title order={4}>{handleSetTitle()}</Title>
				<Group>
					<ToggleTheme />
					<SettingsMenu />
				</Group>
			</Group>

			{messagesListed}
		</Stack>
	);

	return content;
};

export default MainPanel;
