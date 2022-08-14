import type { NextPage } from "next";
import { Paper, Stack, Group, Text, Sx, useMantineColorScheme, Divider, Box } from "@mantine/core";
import { MessageType } from "@/features/types/chat";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { formatRelative } from "date-fns";
import { Eye, Calendar } from "tabler-icons-react";

interface MessagesProps {
	message: MessageType;
	nextMessageDate: Date;
}

const DAY_IN_MILI = 86400000;

const Messages: NextPage<MessagesProps> = ({ message, nextMessageDate }) => {
	const { colorScheme } = useMantineColorScheme();
	const dark = colorScheme === "dark";
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);

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
		<Stack p="sm" mt="xs">
			<Group sx={{ justifyContent: messageAuthor ? "flex-end" : "flex-start" }}>
				<Paper sx={dark ? darkSx : lightSx} shadow="md" p="sm" radius={"lg"}>
					<Stack spacing="sm">
						<Text size="sm">{message.Body}</Text>
						<Group position="apart">
							<Eye
								size={16}
								color={message.Seen?.includes(user!.id) ? "#05a3f4" : "#797c7b"}
							/>
							<Text size="xs">
								{/* {formatRelative(new Date(message.base.createdAt), new Date())} */}
								{message.base.createdAt}
							</Text>
						</Group>
					</Stack>
				</Paper>
			</Group>

			{nextMessageDate.getTime() - new Date(message.base.createdAt).getTime() >
				DAY_IN_MILI && (
				<Divider
					variant="dashed"
					size="sm"
					labelPosition="center"
					label={
						<>
							<Calendar size={20} />
							<Box ml={5}>{new Date(message.base.createdAt).toDateString()}</Box>
						</>
					}
				/>
			)}
		</Stack>
	);
};

export default Messages;
