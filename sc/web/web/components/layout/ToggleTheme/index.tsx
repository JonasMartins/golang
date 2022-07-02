import type { NextPage } from "next";
import { ActionIcon, Group, Switch, useMantineColorScheme } from "@mantine/core";
import { Sun, MoonStars, Settings } from "tabler-icons-react";
import { useState } from "react";
import React from "react";

const ToggleTheme: NextPage = () => {
	const { colorScheme, toggleColorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";

	const [checked, setChecked] = useState(false);

	const handleToggleTheme = () => {
		toggleColorScheme();
		setChecked(!checked);
	};

	return (
		<Group>
			<ActionIcon
				variant="default"
				color={dark ? "yellow" : "indigo"}
				onClick={handleToggleTheme}
				title="Toggle color scheme"
			>
				{dark ? <Sun /> : <MoonStars />}
			</ActionIcon>
			<Switch
				size="lg"
				araia-label="change theme"
				checked={checked}
				color={"indigo"}
				onChange={handleToggleTheme}
			/>
		</Group>
	);
};

export default ToggleTheme;
