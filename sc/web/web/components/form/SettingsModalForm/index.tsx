import { RootState } from "@/app";
import { AlterTypes } from "@/components/notifications/alert/Alert";
import { AlertContent, triggerAlert } from "@/features/alert/alertSlice";
import { setOpenSettingsModal } from "@/features/layout/modalSlicer";
import { UploadProfilePicture, useChangeProfilePictureMutation } from "@/generated/graphql";
import { Button, FileInput, Group, Stack } from "@mantine/core";
import { useForm } from "@mantine/form";
import { IconUpload } from "@tabler/icons";
import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

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
		// TODO: the action to show the alert must be on the
		// main panel component

		if (response.data?.changeProfilePicture.errors.length) {
			const typeAlert: AlterTypes = { type: "ERROR" };
			setErrorInput(x => ({
				...x,
				file: response.data?.changeProfilePicture.errors[0].message || "Server Error",
			}));
		} else {
			dispatch(setOpenSettingsModal(false));
			const typeAlert: AlterTypes = { type: "SUCCESS" };
			const alert: AlertContent = {
				message: "Profile successfuly updated",
				title: "Profile Updated",
				type: typeAlert,
			};
			dispatch(triggerAlert(alert));
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
