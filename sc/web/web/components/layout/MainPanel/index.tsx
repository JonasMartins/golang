import SettingsMenu from "@/components/layout/SettingsMenu";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Group, Stack } from "@mantine/core";
import React from "react";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const content = (
		<Stack>
			<Group position="right" m="md">
				<ToggleTheme />
				<SettingsMenu />
			</Group>
		</Stack>
	);

	return content;
};

export default MainPanel;
