import type { NextPage } from "next";
import ToggleTheme from "@/components/layout/ToggleTheme";
import LoginForm from "@/components/form/LoginForm";
import React from "react";
import { Grid, Group, Title, Stack } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";

const Login: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");

	return (
		<Stack justify="flex-start">
			<Group position="right" m="md">
				<ToggleTheme />
			</Group>
			<Stack spacing="md" justify="center" align="stretch">
				<Title align="center">Login</Title>
				<Grid>
					<Grid.Col span={webScreen ? 4 : 1} />
					<Grid.Col span={webScreen ? 4 : 10}>
						<LoginForm />
					</Grid.Col>
					<Grid.Col span={webScreen ? 4 : 1} />
				</Grid>
			</Stack>
		</Stack>
	);
};

export default Login;
