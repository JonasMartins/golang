import React, { useState } from "react";
import { useForm } from "@mantine/form";
import { Stack, FileInput, Group, Button } from "@mantine/core";
import { UploadProfilePicture, useChangeProfilePictureMutation } from "@/generated/graphql";
import { IconUpload } from "@tabler/icons";
import { RootState } from "@/app";
import { useSelector } from "react-redux";
import GeneralMutationsAlert, { AlterTypes } from "@/components/notifications/alert/Alert";
import { useDispatch } from "react-redux";
import { setOpenSettingsModal } from "@/features/layout/modalSlicer";

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
	const dispatch = useDispatch();

	const form = useForm({
		initialValues: {
			userId: "",
			file: null,
		},
	});

	const HandleUpdateSettings = async (values: UploadProfilePicture) => {
		if (!user) return;

		values.userId = user.id;

		const response = await updatingSettings({
			input: values,
		});

		// TODO HANDLER HERE MAX FILE SIZE
		// dispatch(setOpenSettingsModal(false));
		// TODO: the action to show the alert must be on the
		// main panel component

		if (response.data?.changeProfilePicture.errors.length) {
			const typeAlert: AlterTypes = { type: "ERROR" };
			setErrorInput(x => ({
				...x,
				file: response.data?.changeProfilePicture.errors[0].message || "Server Error",
			}));
			<GeneralMutationsAlert
				title="Error"
				message="Error uploading your new image profile"
				type={typeAlert}
			/>;
		} else {
			const typeAlert: AlterTypes = { type: "SUCCESS" };
			<GeneralMutationsAlert
				title="Updated"
				message="Profile Successfully updated"
				type={typeAlert}
			/>;
		}
	};

	return (
		<Stack mb="sm" mr="xs" ml="xs">
			<form onSubmit={form.onSubmit(values => HandleUpdateSettings(values))}>
				<FileInput
					placeholder="Pick file"
					{...form.getInputProps("file")}
					label="Change Picture"
					icon={<IconUpload size={14} />}
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
