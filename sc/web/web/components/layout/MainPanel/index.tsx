import { Stack, Group } from "@mantine/core";
import React, { useEffect } from "react";
import ToggleTheme from "@/components/layout/ToggleTheme";
import SettingsMenu from "@/components/layout/SettingsMenu";
// import { useGetUsersChatsQuery } from "@/generated/graphql"
import { useUser } from "@/utils/hooks";

interface MainPanelProps {}

const MainPanel: React.FC<MainPanelProps> = () => {
	const user = useUser();

	useEffect(() => {
		if (user) {
			console.log("user ", user);
		}
	}, [user]);

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
