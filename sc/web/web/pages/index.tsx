import type { NextPage } from "next";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Stack, Title } from "@mantine/core";

const Home: NextPage = () => {
	return (
		<Stack>
			<ToggleTheme />
			<Title order={2}>Welcome</Title>
		</Stack>
	);
};

export default Home;
