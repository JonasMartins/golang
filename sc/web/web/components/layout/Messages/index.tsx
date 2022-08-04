import type { NextPage } from "next";
import { Paper, Stack, Group, Text, Sx, useMantineColorScheme } from "@mantine/core";
import { MessageType } from "@/features/types/chat";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { formatRelative } from "date-fns";
import { Eye } from "tabler-icons-react";

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
			backgroundColor: messageAuthor ? "#99e1ba" : "#dde0e2",
		},
	};

	const darkSx: Sx = {
		cursor: "pointer",
		backgroundColor: messageAuthor ? "#07666c" : "#1A1B1E",
		"&:hover": {
			backgroundColor: messageAuthor ? "#075257" : "#070707",
		},
	};

	return (
		<Group sx={{ justifyContent: messageAuthor ? "flex-end" : "flex-start" }}>
			<Paper sx={dark ? darkSx : lightSx} shadow="md" p="sm" radius={"lg"}>
				<Stack spacing="sm">
					<Text>{message.Body}</Text>
					<Group position="apart">
						<Eye size={16} color="#05a3f4" />
						<Text size="xs">
							{formatRelative(new Date(message.base.createdAt), new Date())}
						</Text>
					</Group>
				</Stack>
			</Paper>
		</Group>
	);
};

export default Messages;
