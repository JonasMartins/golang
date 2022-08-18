import type { NextPage } from "next";
import { ActionIcon, Group, Switch, useMantineColorScheme } from "@mantine/core";
import { IconSun, IconMoonStars } from "@tabler/icons";
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
				onClick={handleToggleTheme}
				title="Toggle color scheme"
			>
				{dark ? <IconSun /> : <IconMoonStars />}
			</ActionIcon>
			<Switch
				size="lg"
				araia-label="change theme"
				checked={checked}
				color={"cyan"}
				onChange={handleToggleTheme}
			/>
		</Group>
	);
};

export default ToggleTheme;
