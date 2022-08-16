import React, { useState } from "react";
import { useForm } from "@mantine/form";
import { Stack, FileInput, Group, Button } from "@mantine/core";
import { UploadProfilePicture, useChangeProfilePictureMutation } from "@/generated/graphql";
import { Upload } from "tabler-icons-react";
import { RootState } from "@/app";
import { useSelector } from "react-redux";

interface SettingsModalFormProps {}

type input = {
	file: File | null;
};

type inputError = {
	file: string;
};

const SettingsModalForm: React.FC<SettingsModalFormProps> = ({}) => {
	const user = useSelector((state: RootState) => state.persistedReducer.user.value);
	const [errorInput, setErrorInput] = useState<inputError>({
		file: "",
	});
	const [{}, updatingSettings] = useChangeProfilePictureMutation();

	const form = useForm({
		initialValues: {
			userId: "",
			file: null,
		},
	});

	const HandleUpdateSettings = async (values: UploadProfilePicture) => {
		console.log(values);
		if (!user) return;

		values.userId = user.id;

		const response = await updatingSettings({
			input: values,
		});

		console.log("response ", response);

		if (response.data?.changeProfilePicture.errors.length) {
			setErrorInput(x => ({
				...x,
				file: response.data?.changeProfilePicture.errors[0].message || "Server Error",
			}));
		}
	};

	return (
		<Stack mb="sm" mr="xs" ml="xs">
			<form onSubmit={form.onSubmit(values => HandleUpdateSettings(values))}>
				<FileInput
					placeholder="Pick file"
					{...form.getInputProps("file")}
					label="Change Picture"
					icon={<Upload size={14} />}
					error={errorInput.file}
					accept="image/png,image/jpeg"
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
