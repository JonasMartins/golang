import { Stack, Title } from "@mantine/core";
import React from "react";

interface SideBarProps {}

const SideBar: React.FC<SideBarProps> = () => {
	return (
		<Stack
			sx={theme => ({
				backgroundColor:
					theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2],
				height: "100vh",
			})}
		>
			<Title order={2}>Welcome</Title>
		</Stack>
	);
};

export default SideBar;
