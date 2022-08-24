/* eslint-disable react-hooks/exhaustive-deps */
import NavbarLayout from "@/components/layout/Navbar";
import { setChats as setChatsFromRedux } from "@/features/chat/chatSlicer";
import { setLoggedUser } from "@/features/user/userSlice";
import { GetUsersChatsDocument, GetUsersChatsQuery } from "@/generated/graphql";
import { useUser } from "@/utils/hooks";
import {
	AppShell,
	// Aside,
	Footer,
	Header,
	// MediaQuery,
	Navbar,
	ScrollArea,
	BackgroundImage,
	useMantineTheme,
	useMantineColorScheme,
	Box,
} from "@mantine/core";
import React, { useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { useQuery } from "urql";
import Loader from "@/components/layout/Loader";
import SideBar from "@/components/layout/SideBar";
import Chat from "@/components/layout/Chat";
import CreateMessageForm from "@/components/form/CreateMessage";

interface DashboardProps {}

const Dashboard: React.FC<DashboardProps> = () => {
	const lightBg = "images/light-background.jpg";
	const darkBg = "images/dark-background.jpg";
	const theme = useMantineTheme();
	const [opened, setOpened] = useState(false);
	const user = useUser();
	const dispatch = useDispatch();
	const [userId, setUserId] = useState<string>("");
	const [loadEffect, setLoadEffect] = useState(false);
	const [result, fetch] = useQuery<GetUsersChatsQuery>({
		query: GetUsersChatsDocument,
		pause: true,
		variables: {
			userId,
		},
		requestPolicy: "cache-and-network",
	});

	const { colorScheme } = useMantineColorScheme();
	const handleGetData = useCallback(() => {
		if (userId.length) {
			fetch();
		}
	}, [userId, fetch]);

	useEffect(() => {
		setLoadEffect(true);
		if (!user) return;

		if (user.id) {
			dispatch(setLoggedUser(user));
			setUserId(user.id);
		}

		handleGetData();

		setTimeout(() => {
			setLoadEffect(false);
		}, 500);
	}, [user]);

	useEffect(() => {
		if (result.data?.getUsersChats.chats.length && user) {
			dispatch(setChatsFromRedux(result.data.getUsersChats.chats));
		}
	}, [result.fetching, result.data?.getUsersChats.chats.length, dispatch]);

	return !user || loadEffect || result.fetching ? (
		<Loader />
	) : (
		<AppShell
			styles={{
				main: {
					background:
						theme.colorScheme === "dark" ? theme.colors.dark[8] : theme.colors.gray[0],
				},
			}}
			navbarOffsetBreakpoint="sm"
			asideOffsetBreakpoint="sm"
			navbar={
				<Navbar p="md" hiddenBreakpoint="sm" hidden={!opened} width={{ sm: 300, lg: 500 }}>
					<Navbar.Section grow component={ScrollArea} mx="-xs" px="xs">
						<SideBar loggedUser={user} />
					</Navbar.Section>
				</Navbar>
			}
			/*
			aside={
				<MediaQuery smallerThan="sm" styles={{ display: "none" }}>
					<Aside p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
						<Text>Application sidebar</Text>
					</Aside>
				</MediaQuery>
			} */
			footer={
				<Footer height={110} p="xs">
					<CreateMessageForm />
				</Footer>
			}
			header={
				<Header height={60} p="md">
					<NavbarLayout opened={opened} setOpened={setOpened} />
				</Header>
			}
		>
			<BackgroundImage src={colorScheme === "dark" ? darkBg : lightBg}>
				<Chat />
			</BackgroundImage>
		</AppShell>
	);
};

export default Dashboard;
