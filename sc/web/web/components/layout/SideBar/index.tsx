import SettingsMenu from "@/components/layout/SettingsMenu";
import { GetUsersChatsQuery } from "@/generated/graphql";
import { Grid, Input, Stack } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import React from "react";
import { Search } from "tabler-icons-react";
import ChatsSideBar from "@/components/layout/ChatsSideBar";
import { UserJwt } from "@/utils/hooks";

interface SideBarProps {
	chats: GetUsersChatsQuery | undefined;
	loggedUser: UserJwt;
}

const SideBar: React.FC<SideBarProps> = ({ chats, loggedUser }) => {
	const webScreen = useMediaQuery("(min-width: 900px)");

	const chatsEle = (
		<Stack
			sx={() => ({
				flexDirection: "column",
				alignItems: "stretch",
				justifyContent: "center",
				flexGrow: 1,
			})}
			spacing="lg"
			mt="lg"
		>
			{chats?.getUsersChats.chats.map(x => (
				<ChatsSideBar key={x.base.id} chat={x} currentUserId={loggedUser.id} />
			))}
		</Stack>
	);

	return (
		<Stack p="md">
			<Grid gutter="sm" align={"center"} grow={true}>
				<Grid.Col span={webScreen ? 12 : 11}>
					<Input icon={<Search />} variant="filled" placeholder="Search" />
				</Grid.Col>
				{!webScreen && (
					<Grid.Col span={1}>
						<SettingsMenu />
					</Grid.Col>
				)}
				{chatsEle}
			</Grid>
		</Stack>
	);
};

export default SideBar;
