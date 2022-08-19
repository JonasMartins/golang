import { RootState } from "@/app";
import { addMessage } from "@/features/chat/chatSlicer";
import { MessageType } from "@/features/types/chat";
import { CreateMessageInput, useCreateMessageMutation } from "@/generated/graphql";
import { uuidv4Like } from "@/utils/aux/chat.aux";
import { ActionIcon, Group, Stack, Textarea } from "@mantine/core";
import { useForm } from "@mantine/form";
import type { NextPage } from "next";
import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { IconMoodSmile, IconSend } from "@tabler/icons";

interface CreateMessageFormProps {}

type input = {
	body: string;
};

const CreateMessageForm: NextPage<CreateMessageFormProps> = ({}) => {
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);
	const chatFocused = useSelector((state: RootState) => state.persistedReducer.chat.value);
	const [errorInput, setErrorInput] = useState<input>({
		body: "",
	});
	const dispatch = useDispatch();
	const [{}, createMessage] = useCreateMessageMutation();

	const form = useForm({
		initialValues: {
			body: "",
			authorId: "",
			chatId: "",
		},
	});

	const HandleCreateMessage = async (values: CreateMessageInput) => {
		if (!handleIfMessageBodyHasValidCharacters()) return;
		if (!user || !chatFocused) {
			setErrorInput(prevState => ({
				...prevState,
				body: "Server error",
			}));
			return;
		}

		values.authorId = user.id;
		values.chatId = chatFocused.base.id;

		const response = await createMessage({
			input: values,
		});

		if (response.data?.createMessage.errors.length) {
			setErrorInput(prevState => ({
				...prevState,
				body: response.data?.createMessage.errors[0].message || "Server Error",
			}));
		} else {
			const newMessage: MessageType = {
				Author: {
					name: user.name,
				},
				ChatId: chatFocused.base.id,
				base: {
					id: uuidv4Like(),
					createdAt: new Date().getTime(),
				},
				Body: values.body,
			};
			form.setValues({
				body: "",
				authorId: user.id,
				chatId: chatFocused.base.id,
			});
			dispatch(addMessage(newMessage));
		}
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
					if (handleIfInputCanSubmitOnEnter(e.key)) {
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
							disabled={chatFocused || errorInput.body ? false : true}
							sx={{ flexGrow: 1 }}
							p="sm"
							radius="lg"
							size="sm"
							{...form.getInputProps("body")}
							required
							rightSection={
								<ActionIcon
									size="xl"
									radius="xl"
									mr="lg"
									disabled={chatFocused || errorInput.body ? false : true}
								>
									<IconMoodSmile />
								</ActionIcon>
							}
							error={errorInput.body}
						/>
					</Group>

					<ActionIcon size="xl" radius="xl" color="cyan" component="button" type="submit">
						<IconSend />
					</ActionIcon>
				</Group>
			</form>
		</Stack>
	);
};

export default CreateMessageForm;
