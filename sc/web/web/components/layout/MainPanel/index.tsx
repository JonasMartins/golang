import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Group, Stack, Text } from "@mantine/core";
import React from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/Redux";
import { GetChatTitle } from "@/utils/aux/chat.aux";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const chatFocused = useSelector((state: RootState) => state.chat.value);
	const user = useSelector((state: RootState) => state.user.value);

	const handleSetTitle = (): string => {
		if (user && chatFocused) {
			return GetChatTitle(chatFocused, user.id);
		}
		return "Unknown";
	};

	const content = (
		<Stack>
			<Group position="apart" m="md">
				<Text>{handleSetTitle()}</Text>
				<Group>
					<ToggleTheme />
					<SettingsMenu />
				</Group>
			</Group>
		</Stack>
	);

	return content;
};

export default MainPanel;
