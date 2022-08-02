import MainPanel from "@/components/layout/MainPanel";
import SideBar from "@/components/layout/SideBar";
import { Grid } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
import { useUser } from "@/utils/hooks";
import React, { useCallback, useEffect, useState } from "react";
import Loader from "@/components/layout/Loader";
import { GetUsersChatsDocument, GetUsersChatsQuery } from "@/generated/graphql";
import { useQuery } from "urql";

const Dashboard: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");
	const user = useUser();
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

	const handleGetData = useCallback(() => {
		if (userId.length) {
			fetch();
		}
	}, [userId, fetch]);

	useEffect(() => {
		setLoadEffect(true);
		if (!user) return;

		if (user.id) {
			setUserId(user.id);
		}

		handleGetData();

		setTimeout(() => {
			setLoadEffect(false);
		}, 500);
	}, [user, handleGetData]);

	const web =
		!user || loadEffect ? (
			<Loader />
		) : (
			<Grid>
				<Grid.Col span={4}>
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
			<Grid>
				<Grid.Col span={12}>
					<SideBar chats={result.data} loggedUser={user} />
				</Grid.Col>
			</Grid>
		);

	return webScreen ? web : mobile;
};

export default Dashboard;
