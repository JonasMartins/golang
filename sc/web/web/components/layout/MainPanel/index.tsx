import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Group, Stack, Text, useMantineColorScheme } from "@mantine/core";
import React from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/Redux";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import MessageComp from "@/components/layout/Messages";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.chat.value);
	const user = useSelector((state: RootState) => state.user.value);
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";
	const gradientLight = ["#c8cbc9", "#ffffff"];

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
			{chatFocused?.Messages.map(x => (
				<MessageComp key={x.base.createdAt} message={x} />
			))}
		</Stack>
	);

	const content = (
		<Stack
			sx={theme => ({
				backgroundImage: dark
					? theme.fn.linearGradient(133, "#272d28", "#3f4540")
					: theme.fn.linearGradient(133, "#c8cbc9", "#ffffff"),
				height: "100vh",
			})}
		>
			<Group position="apart" m="md">
				<Text>{handleSetTitle()}</Text>
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
