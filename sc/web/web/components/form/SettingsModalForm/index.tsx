import React from "react";
import { useForm } from "@mantine/form";
import { Stack, FileInput, Group, Button } from "@mantine/core";
import { UploadProfilePicture } from "@/generated/graphql";
import { Upload } from "tabler-icons-react";
interface SettingsModalFormProps {}

type input = {
	file: File;
};

const SettingsModalForm: React.FC<SettingsModalFormProps> = ({}) => {
	const form = useForm({
		initialValues: {
			userId: "",
			file: null,
		},
	});

	const HandleUpdateSettings = (values: UploadProfilePicture) => {
		console.log(values);
	};

	return (
		<Stack mb="sm" mr="xs" ml="xs">
			<form onSubmit={form.onSubmit(values => HandleUpdateSettings(values))}>
				<FileInput
					placeholder="Pick file"
					{...form.getInputProps("file")}
					label="Change Picture"
					icon={<Upload size={14} />}
				/>
				<Group grow={true} mt="md">
					<Button
						variant="gradient"
						gradient={{ from: "indigo", to: "cyan" }}
						type="submit"
					>
						Update
					</Button>
				</Group>
			</form>
		</Stack>
	);
};

export default SettingsModalForm;
