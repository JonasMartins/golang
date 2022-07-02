import type { NextPage } from "next";
import ToggleTheme from "@/components/layout/ToggleTheme";
import LoginForm from "@/components/form/LoginForm";
import React from "react";
import { Box, Text, Group, Title } from "@mantine/core";

const Login: NextPage = () => {
	return (
		<React.Fragment>
			<Box>
				<Group position="right" m="md">
					<ToggleTheme />
				</Group>
				<Title align="center">Login</Title>

				<LoginForm />
			</Box>
		</React.Fragment>
	);
};

export default Login;
