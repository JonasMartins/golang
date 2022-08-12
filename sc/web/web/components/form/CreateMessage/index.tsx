import type { NextPage } from "next";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { useForm } from "@mantine/form";
import { Textarea, ActionIcon, Button, Stack, Grid, Group } from "@mantine/core";
import { CreateMessageInput } from "@/generated/graphql";
import { useDispatch } from "react-redux";
import { MessageType } from "@/features/types/chat";
import { addMessage } from "@/features/chat/chatSlicer";
import { MoodSmile, Send } from "tabler-icons-react";

interface CreateMessageFormProps {
	showSubmitButton: boolean;
}

type input = {
	body: string;
	authorId: string;
	chatId: string;
};

const CreateMessageForm: NextPage<CreateMessageFormProps> = ({ showSubmitButton }) => {
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const dispatch = useDispatch();

	const form = useForm({
		initialValues: {
			body: "",
			authorId: "",
			chatId: "",
		},
	});

	const HandleCreateMessage = (values: CreateMessageInput) => {
		const newMessage: MessageType = {
			Author: {
				name: user?.name!,
			},
			ChatId: chatFocused?.base.id || "",
			base: {
				id: "",
				createdAt: new Date().toDateString(),
			},
			Body: values.body,
		};
		form.setValues({
			body: "",
			authorId: user?.id || "",
			chatId: chatFocused?.base.id || "",
		});
		dispatch(addMessage(newMessage));
	};
	/**
	 * If a enter has been clicked, this method will check
	 * if the message body has new lines, indicating that the user
	 * has a more elaborated text, if so, then will return false
	 * openning a button on side allowing the user to put how many
	 * new lines he wants it on the message body
	 *
	 * @param key string
	 * @returns boolean
	 */
	const handleIfInputCanSubmitOnEnter = (key: string): boolean => {
		if (key === "Enter") {
			return true;
		}
		return false;
	};

	/**
	 * If the message has only new lines, tabs of blanck spaces
	 * then it will return false, not allowing the message to be created
	 * @returns boolean
	 */
	const handleIfMessageBodyHasValidCharacters = (): boolean => {
		return /\S/.test(form.values.body);
	};

	return (
		<Stack mb="sm" mr="xs" ml="xs">
			<form
				onKeyDown={e => {
					if (handleIfInputCanSubmitOnEnter(e.key) && !showSubmitButton) {
						if (handleIfMessageBodyHasValidCharacters()) {
							HandleCreateMessage(form.values);
						}
					}
				}}
				onSubmit={form.onSubmit(values => HandleCreateMessage(values))}
			>
				<Group spacing={0}>
					<Group sx={{ flexGrow: 1 }}>
						<Textarea
							sx={{ flexGrow: 1 }}
							p="sm"
							radius="lg"
							size="sm"
							{...form.getInputProps("body")}
							required
							rightSection={
								<ActionIcon size="xl" radius="xl" mr="lg">
									<MoodSmile />
								</ActionIcon>
							}
						/>
					</Group>
					{showSubmitButton && (
						<ActionIcon
							size="xl"
							radius="xl"
							color="cyan"
							component="button"
							type="submit"
						>
							<Send />
						</ActionIcon>
					)}
				</Group>
			</form>
		</Stack>
	);
};

export default CreateMessageForm;
