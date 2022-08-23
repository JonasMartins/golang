import NavbarLayout from "@/components/layout/Navbar";
import {
	AppShell,
	Aside,
	Footer,
	Header,
	MediaQuery,
	Navbar,
	Text,
	useMantineTheme,
} from "@mantine/core";
import { useState } from "react";

interface AppShellDemoProps {}

const AppShellDemo: React.FC<AppShellDemoProps> = () => {
	const theme = useMantineTheme();
	const [opened, setOpened] = useState(false);
	return (
		<AppShell
			styles={{
				main: {
					background:
						theme.colorScheme === "dark" ? theme.colors.dark[8] : theme.colors.gray[0],
				},
			}}
			navbarOffsetBreakpoint="sm"
			asideOffsetBreakpoint="sm"
			navbar={
				<Navbar p="md" hiddenBreakpoint="sm" hidden={!opened} width={{ sm: 200, lg: 300 }}>
					<Text>Application navbar</Text>
				</Navbar>
			}
			aside={
				<MediaQuery smallerThan="sm" styles={{ display: "none" }}>
					<Aside p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
						<Text>Application sidebar</Text>
					</Aside>
				</MediaQuery>
			}
			footer={
				<Footer height={60} p="md">
					Application footer
				</Footer>
			}
			header={
				<Header height={70} p="md">
					<NavbarLayout opened={opened} setOpened={setOpened} />
				</Header>
			}
		>
			<Text>Resize app to see responsive navbar in action</Text>
		</AppShell>
	);
};

export default AppShellDemo;
