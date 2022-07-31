import { Stack, Group } from "@mantine/core";
import React from "react";
import ToggleTheme from "@/components/layout/ToggleTheme";
import SettingsMenu from "@/components/layout/SettingsMenu";
// import { useGetUsersChatsQuery } from "@/generated/graphql"

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	return (
		<Stack>
			<Group position="right" m="md">
				<ToggleTheme />
				<SettingsMenu />
			</Group>
		</Stack>
	);
};

export default MainPanel;
