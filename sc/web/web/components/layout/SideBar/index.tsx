import { RootState } from "@/app";
import ChatsSideBar from "@/components/layout/ChatsSideBar";
import { updadeChatsFromCommingNewMessageSubscription } from "@/features/chat/chatSlicer";
import { ChatType } from "@/features/types/chat";
import { useMessageSendedSubscription } from "@/generated/graphql";
import { GetChatTitle } from "@/utils/aux/chat.aux";
import { UserJwt } from "@/utils/hooks";
import { Grid, Input, Stack } from "@mantine/core";
import { IconSearch } from "@tabler/icons";
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

interface SideBarProps {
	loggedUser: UserJwt;
}

const SideBar: React.FC<SideBarProps> = ({ loggedUser }) => {
	const dispatch = useDispatch();
	const [searchTerm, setSearchTerm] = useState("");
	const chatsFromReducer = useSelector((state: RootState) => state.persistedReducer.chat.chats);
	const [chatsState, setChatsState] = useState<ChatType[] | undefined>(chatsFromReducer);
	const [resultNewMessage] = useMessageSendedSubscription();

	useEffect(() => {
		if (resultNewMessage.data?.messageSended) {
			if (loggedUser.id !== resultNewMessage.data.messageSended.AuthorId) {
				dispatch(
					updadeChatsFromCommingNewMessageSubscription(
						resultNewMessage.data.messageSended
					)
				);
			}
		}
	}, [resultNewMessage, dispatch, loggedUser]);

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
				<Grid.Col span={12}>
					<Input
						icon={<IconSearch />}
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

				{chatsEle}
			</Grid>
		</Stack>
	);
};

export default SideBar;
