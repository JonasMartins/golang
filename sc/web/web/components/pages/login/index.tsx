import LoginForm from "@/components/form/LoginForm";
import RegisterForm from "@/components/form/RegisterForm";
import ToggleTheme from "@/components/layout/ToggleTheme";
import { Grid, Group, Stack, Title, Tabs } from "@mantine/core";
import { useMediaQuery } from "@mantine/hooks";
import type { NextPage } from "next";
import { useState } from "react";

const Login: NextPage = () => {
	const webScreen = useMediaQuery("(min-width: 900px)");
	const [activeTab, setActiveTab] = useState<string | null>("Login");
	return (
		<Stack justify="flex-start">
			<Group position="right" m="md"></Group>
			<Stack spacing="md" justify="center" align="stretch">
				<Title align="center">{activeTab}</Title>
				<Grid>
					<Grid.Col span={webScreen ? 4 : 1} />
					<Grid.Col span={webScreen ? 4 : 10}>
						<Tabs defaultValue={"Login"} onTabChange={setActiveTab}>
							<Tabs.List grow>
								<Tabs.Tab value="Login" aria-label="login">
									Login
								</Tabs.Tab>
								<Tabs.Tab value="Register" aria-label="register">
									Register
								</Tabs.Tab>
							</Tabs.List>
							<Tabs.Panel value="Login">
								<LoginForm />
							</Tabs.Panel>
							<Tabs.Panel value="Register">
								<RegisterForm />
							</Tabs.Panel>
						</Tabs>
					</Grid.Col>
					<Grid.Col span={webScreen ? 4 : 1} />
				</Grid>
			</Stack>
		</Stack>
	);
};

export default Login;
