import { Stack, Group } from "@mantine/core";
import React from "react";
import ToggleTheme from "@/components/layout/ToggleTheme";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	return (
		<Stack>
			<Group position="right" m="md">
				<ToggleTheme />
			</Group>
		</Stack>
	);
};

export default MainPanel;
