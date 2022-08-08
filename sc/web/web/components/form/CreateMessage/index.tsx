import type { NextPage } from "next";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { useForm } from "@mantine/form";
import { Textarea, ActionIcon, Button, Stack } from "@mantine/core";
import { CreateMessageInput } from "@/generated/graphql";
import { useDispatch } from "react-redux";
import { MessageType } from "@/features/types/chat";
import { addMessage } from "@/features/chat/chatSlicer";

interface CreateMessageFormProps {}

type input = {
	body: string;
	authorId: string;
	chatId: string;
};

const CreateMessageForm: NextPage<CreateMessageFormProps> = () => {
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);
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
			base: {
				createdAt: new Date(),
			},
			Body: values.body,
		};
		dispatch(addMessage(newMessage));
	};

	return (
		<Stack mb="sm" mr="xs" ml="xs">
			<form onSubmit={form.onSubmit(values => HandleCreateMessage(values))}>
				<Textarea
					p="sm"
					radius="lg"
					size="sm"
					{...form.getInputProps("body")}
					required
					rightSection={<Button type="submit">Send</Button>}
				/>
			</form>
		</Stack>
	);
};

export default CreateMessageForm;
