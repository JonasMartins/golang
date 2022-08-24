import SettingsMenu from "@/components/layout/SettingsMenu";
import { Logo } from "@/components/mantine/examples/_logo";
import {
	ActionIcon,
	Burger,
	Group,
	MediaQuery,
	useMantineColorScheme,
	useMantineTheme,
	Text,
} from "@mantine/core";
import { IconMoonStars, IconSun } from "@tabler/icons";
import { Dispatch, SetStateAction, useEffect, useState } from "react";

interface NavbarProps {
	opened: boolean;
	setOpened: Dispatch<SetStateAction<boolean>>;
}

const Navbar: React.FC<NavbarProps> = ({ opened, setOpened }) => {
	const theme = useMantineTheme();
	const [sizeScreen, getSizeScreen] = useState({
		dynamicWidth: window.innerWidth,
		dynamicHeight: window.innerHeight,
	});

	const { colorScheme, toggleColorScheme } = useMantineColorScheme();

	const setDimension = () => {
		getSizeScreen({
			dynamicHeight: window.innerHeight,
			dynamicWidth: window.innerWidth,
		});
	};

	useEffect(() => {
		window.addEventListener("resize", setDimension);
		return () => {
			window.removeEventListener("resize", setDimension);
		};
	}, [sizeScreen]);

	return (
		<div style={{ display: "flex", alignItems: "center", height: "100%" }}>
			<MediaQuery largerThan="sm" styles={{ display: "none" }}>
				<Burger
					opened={opened}
					onClick={() => setOpened(o => !o)}
					size="sm"
					color={theme.colors.gray[6]}
					mr="xl"
				/>
			</MediaQuery>
			<Group sx={{ height: "100%", width: "100%" }} px={20} position="apart" grow>
				<Group>
					<Logo colorScheme={colorScheme} />
					<ActionIcon variant="default" onClick={() => toggleColorScheme()} size={30}>
						{colorScheme === "dark" ? (
							<IconSun size={16} />
						) : (
							<IconMoonStars size={16} />
						)}
					</ActionIcon>
					<Text>{sizeScreen.dynamicHeight}</Text>
				</Group>
				<SettingsMenu />
			</Group>
		</div>
	);
};

export default Navbar;