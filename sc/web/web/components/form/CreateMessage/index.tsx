import type { NextPage } from "next";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { useForm } from "@mantine/form";
import { Textarea, ActionIcon, Button } from "@mantine/core";
import { MoodSmile } from "tabler-icons-react";
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
	const user = useSelector((state: RootState) => state.user.value);
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
		<>
			<form onSubmit={form.onSubmit(values => HandleCreateMessage(values))}>
				<Textarea
					p="sm"
					radius="lg"
					size="sm"
					required
					rightSection={
						<ActionIcon>
							{/* <MoodSmile /> */}
							<Button
								variant="gradient"
								gradient={{ from: "indigo", to: "cyan" }}
								type="submit"
							>
								Send
							</Button>
						</ActionIcon>
					}
				/>
			</form>
		</>
	);
};

export default CreateMessageForm;
