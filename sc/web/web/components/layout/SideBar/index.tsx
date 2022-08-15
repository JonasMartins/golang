import { RootState } from "@/app";
import ChatsSideBar from "@/components/layout/ChatsSideBar";
import SettingsMenu from "@/components/layout/SettingsMenu";
import { ChatType } from "@/features/types/chat";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import { UserJwt } from "@/utils/hooks";
import { Grid, Input, Stack } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import React, { useState } from "react";
import { useSelector } from "react-redux";
import { Search } from "tabler-icons-react";

interface SideBarProps {
	loggedUser: UserJwt;
}

const SideBar: React.FC<SideBarProps> = ({ loggedUser }) => {
	const webScreen = useMediaQuery("(min-width: 900px)");
	const [searchTerm, setSearchTerm] = useState("");
	const chatsFromReducer = useSelector((state: RootState) => state.persistedReducer.chat.chats);
	const [chatsState, setChatsState] = useState<ChatType[] | undefined>(chatsFromReducer);

	const chatsEle = (
		<Stack
			sx={() => ({
				flexDirection: "column",
				alignItems: "stretch",
				justifyContent: "center",
				flexGrow: 1,
			})}
			spacing="xs"
			mt="xs"
		>
			{chatsState &&
				chatsState.map(x => (
					<ChatsSideBar key={x.base.id} chat={x} title={GetChatTitle(x, loggedUser.id)} />
				))}
		</Stack>
	);

	return (
		<Stack p="md">
			<Grid gutter="sm" align={"center"} grow={true}>
				<Grid.Col span={webScreen ? 12 : 11}>
					<Input
						icon={<Search />}
						value={searchTerm}
						onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
							setSearchTerm(e.target.value);

							if (e.target.value.length >= 2) {
								let regexTerm = "[ˆ,]*" + e.target.value.toLowerCase() + "[,$]*";

								setChatsState(prevChats => {
									return prevChats?.filter(x =>
										GetChatTitle(x, loggedUser.id)
											.toLowerCase()
											.match(regexTerm)
									);
								});
							} else if (e.target.value.length === 0) {
								setChatsState(chatsFromReducer);
							}
						}}
						variant="filled"
						placeholder="Search a chat"
					/>
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
