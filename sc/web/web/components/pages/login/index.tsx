import LoginForm from "@/components/form/LoginForm";
import RegisterForm from "@/components/form/RegisterForm";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Grid, Group, Stack, Title, Tabs } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
import { useState } from "react";

const Login: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");
	const [activeTab, setActiveTab] = useState(0);
	return (
		<Stack justify="flex-start">
			<Group position="right" m="md">
				<ToggleTheme />
			</Group>
			<Stack spacing="md" justify="center" align="stretch">
				<Title align="center">{activeTab === 0 ? "Login" : "Register"}</Title>
				<Grid>
					<Grid.Col span={webScreen ? 4 : 1} />
					<Grid.Col span={webScreen ? 4 : 10}>
						<Tabs grow active={activeTab} onTabChange={setActiveTab}>
							<Tabs.Tab label="Login" aria-label="login">
								<LoginForm />
							</Tabs.Tab>
							<Tabs.Tab label="Register" aria-label="register">
								<RegisterForm />
							</Tabs.Tab>
						</Tabs>
					</Grid.Col>
					<Grid.Col span={webScreen ? 4 : 1} />
				</Grid>
			</Stack>
		</Stack>
	);
};

export default Login;
