import MainPanel from "@/components/layout/MainPanel";
import SideBar from "@/components/layout/SideBar";
import { Grid } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
import { useUser } from "@/utils/hooks";
import React, { useEffect, useState } from "react";
import Loader from "@/components/layout/Loader";

const Dashboard: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");

	const user = useUser();
	const [loadEffect, setLoadEffect] = useState(false);

	useEffect(() => {
		setLoadEffect(true);
		if (!user) return;
		setTimeout(() => {
			setLoadEffect(false);
		}, 500);
	}, [user]);

	const web =
		!user || loadEffect ? (
			<Loader />
		) : (
			<Grid>
				<Grid.Col span={4}>
					<SideBar />
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
					<SideBar />
				</Grid.Col>
			</Grid>
		);

	return webScreen ? web : mobile;
};

export default Dashboard;
