import type { NextPage } from "next";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import { useForm } from "@mantine/form";
import { Textarea, ActionIcon } from "@mantine/core";
import { MoodSmile } from "tabler-icons-react";

interface CreateMessageFormProps {}

type input = {
	body: string;
	authorId: string;
	chatId: string;
};

const CreateMessageForm: NextPage<CreateMessageFormProps> = () => {
	const user = useSelector((state: RootState) => state.user.value);

	const form = useForm({
		initialValues: {
			body: "",
			authorId: "",
			chatId: "",
		},
	});

	return (
		<>
			<form onSubmit={form.onSubmit(values => console.log(values))}>
				<Textarea
					p="sm"
					radius="lg"
					size="lg"
					required
					rightSection={
						<ActionIcon>
							<MoodSmile />
						</ActionIcon>
					}
				/>
			</form>
		</>
	);
};

export default CreateMessageForm;
