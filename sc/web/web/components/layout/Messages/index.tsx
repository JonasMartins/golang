import type { NextPage } from "next";
import { Paper, Stack, Group, Title, Text, Sx, useMantineColorScheme } from "@mantine/core";
import { MessageType } from "@/features/types/chat";
import { RootState } from "@/Redux";
import { useSelector } from "react-redux";
import { formatRelative } from "date-fns";

interface MessagesProps {
	message: MessageType;
}

const Messages: NextPage<MessagesProps> = ({ message }) => {
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";
	const user = useSelector((state: RootState) => state.user.value);

	let messageAuthor = message.Author.name === user?.name;

	const lightSx: Sx = {
		cursor: "pointer",
		backgroundColor: messageAuthor ? "#b2ffcb" : "#fff",
		"&:hover": {
			backgroundColor: messageAuthor ? "#9fe1b4" : "#b2ffcb",
		},
	};

	const darkSx: Sx = {
		cursor: "pointer",
		backgroundColor: messageAuthor ? "#07858d" : "#1A1B1E",
		"&:hover": {
			backgroundColor: messageAuthor ? "#07666c" : "#070707",
		},
	};

	return (
		<Group sx={{ justifyContent: messageAuthor ? "flex-end" : "flex-start" }}>
			<Paper sx={dark ? darkSx : lightSx} shadow="md" p="sm" withBorder radius={"lg"}>
				<Stack spacing="sm">
					<Text>{message.Body}</Text>
					<Text size="xs" align="right">
						{formatRelative(new Date(message.base.createdAt), new Date())}
					</Text>
				</Stack>
			</Paper>
		</Group>
	);
};

export default Messages;
