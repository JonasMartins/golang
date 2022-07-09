import MainPanel from "@/components/layout/MainPanel";
import SideBar from "@/components/layout/SideBar";
import { Grid } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
const Home: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");

	const web = (
		<Grid>
			<Grid.Col span={4}>
				<SideBar />
			</Grid.Col>
			<Grid.Col span={8}>
				<MainPanel />
			</Grid.Col>
		</Grid>
	);

	const mobile = (
		<Grid>
			<Grid.Col span={12}>
				<SideBar />
			</Grid.Col>
		</Grid>
	);

	return webScreen ? web : mobile;
};

export default Home;
