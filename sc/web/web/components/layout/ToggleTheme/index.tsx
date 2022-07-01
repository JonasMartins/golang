import type { NextPage } from "next";
import { ActionIcon, Group, Switch, useMantineColorScheme } from "@mantine/core";
import { Sun, MoonStars } from "tabler-icons-react";
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
				variant="transparent"
				color={dark ? "yellow" : "cyan"}
				onClick={handleToggleTheme}
				title="Toggle color scheme"
			>
				{dark ? <Sun size="lg" /> : <MoonStars size="lg" />}
			</ActionIcon>
			<Switch size="lg" checked={checked} color={"teal"} onChange={handleToggleTheme} />
		</Group>
	);
};

export default ToggleTheme;
