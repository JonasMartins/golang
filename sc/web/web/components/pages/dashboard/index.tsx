/* eslint-disable react-hooks/exhaustive-deps */
import MainPanel from "@/components/layout/MainPanel";
import SideBar from "@/components/layout/SideBar";
import { Grid, useMantineColorScheme } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
import { useUser } from "@/utils/hooks";
import React, { useCallback, useEffect, useState } from "react";
import Loader from "@/components/layout/Loader";
import { GetUsersChatsDocument, GetUsersChatsQuery } from "@/generated/graphql";
import { useQuery } from "urql";
import { setFocusedChat } from "@/features/chat/chatSlicer";
import { setLoggedUser } from "@/features/user/userSlice";
import { useDispatch } from "react-redux";

const Dashboard: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");
	const user = useUser();
	const dispatch = useDispatch();
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";
	const [userId, setUserId] = useState<string>("");
	const [loadEffect, setLoadEffect] = useState(false);
	const [result, fetch] = useQuery<GetUsersChatsQuery>({
		query: GetUsersChatsDocument,
		pause: true,
		variables: {
			userId,
		},
		//requestPolicy: "cache-and-network",
	});

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
	}, [user, handleGetData]);

	useEffect(() => {
		if (result.data?.getUsersChats.chats.length) {
			dispatch(setFocusedChat(result.data.getUsersChats.chats[0]));
		}
	}, [result.fetching, result.data?.getUsersChats.chats.length, dispatch]);

	const web =
		!user || loadEffect ? (
			<Loader />
		) : (
			<Grid
				gutter={0}
				sx={theme => ({
					backgroundColor: dark ? theme.colors.dark[4] : theme.colors.gray[2],
					height: "100vh",
				})}
			>
				<Grid.Col span={4} sx={{ borderRightStyle: "double" }}>
					<SideBar chats={result.data} loggedUser={user} />
				</Grid.Col>
				<Grid.Col span={8}>
					<MainPanel />
				</Grid.Col>
			</Grid>
		);

	const mobile =
		!user || loadEffect ? (
			<Loader />
		) : (
			<Grid
				gutter={0}
				sx={theme => ({
					backgroundColor: dark ? theme.colors.dark[5] : theme.colors.gray[2],
					height: "100vh",
				})}
			>
				<Grid.Col span={12}>
					<SideBar chats={result.data} loggedUser={user} />
				</Grid.Col>
			</Grid>
		);

	return webScreen ? web : mobile;
};

export default Dashboard;
