import type { NextPage } from "next";
import {
	Paper,
	Stack,
	Group,
	Text,
	Sx,
	useMantineColorScheme,
	Divider,
	Box,
	Accordion,
	ActionIcon,
} from "@mantine/core";
import { MessageType } from "@/features/types/chat";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { formatRelative } from "date-fns";
import { IconEye, IconCalendar, IconChevronDown } from "@tabler/icons";

interface MessagesProps {
	message: MessageType;
	nextMessageDate: Date;
	messagesUnSeenIndicator: boolean;
}

const DAY_IN_MILI = 86400000;

const Messages: NextPage<MessagesProps> = ({
	message,
	nextMessageDate,
	messagesUnSeenIndicator,
}) => {
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
		<Stack p="xs" mt="xs">
			{messagesUnSeenIndicator && (
				<Paper radius="lg" p={1} withBorder>
					<Group position="center" align="center">
						<Text size={"sm"}>Unseen Messages</Text>
						<ActionIcon>
							<IconChevronDown size={15} />
						</ActionIcon>
					</Group>
				</Paper>
			)}

			<Group sx={{ justifyContent: messageAuthor ? "flex-end" : "flex-start" }}>
				<Paper sx={dark ? darkSx : lightSx} shadow="md" p="xs" radius={"lg"}>
					<Stack spacing="sm">
						<Text size="xs">{message.Body}</Text>
						<Group position="apart">
							<IconEye
								size={14}
								color={message.Seen?.includes(user!.id) ? "#05a3f4" : "#797c7b"}
							/>
							<Text size="xs">
								{formatRelative(new Date(message.base.createdAt), new Date())}
								{/* {message.base.createdAt} */}
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
							<IconCalendar size={20} />
							<Box ml={5}>{new Date(message.base.createdAt).toDateString()}</Box>
						</>
					}
				/>
			)}
		</Stack>
	);
};

export default Messages;
