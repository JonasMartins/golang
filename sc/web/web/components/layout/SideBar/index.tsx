import { Stack, Title, Input, Group, Grid } from "@mantine/core";
import React, { useEffect } from "react";
import { Search } from "tabler-icons-react";
import { useMediaQuery } from "@mantine/hooks";
import SettingsMenu from "@/components/layout/SettingsMenu";
import { GetUsersChatsQuery } from "@/generated/graphql";

interface SideBarProps {
	chats: GetUsersChatsQuery | undefined;
}

const SideBar: React.FC<SideBarProps> = ({ chats }) => {
	useEffect(() => {
		console.log(chats);
	}, []);

	const webScreen = useMediaQuery("(min-width: 900px)");
	return (
		<Stack
			sx={theme => ({
				backgroundColor:
					theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2],
				height: "100vh",
			})}
			p="sm"
		>
			<Grid gutter="sm" align={"center"} grow={true}>
				<Grid.Col span={webScreen ? 12 : 11}>
					<Input icon={<Search />} variant="filled" placeholder="Search" />
				</Grid.Col>
				{!webScreen && (
					<Grid.Col span={1}>
						<SettingsMenu />
					</Grid.Col>
				)}
			</Grid>
		</Stack>
	);
};

export default SideBar;
