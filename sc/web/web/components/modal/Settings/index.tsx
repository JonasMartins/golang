import { RootState } from "@/app";
import { setOpenSettingsModal } from "@/features/layout/modalSlicer";
import { Avatar, Modal, Stack } from "@mantine/core";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import SettingsModalForm from "@/components/form/SettingsModalForm";

interface SettingsModalProps {}

const SettingsModal: React.FC<SettingsModalProps> = ({}) => {
	const dispatch = useDispatch();
	const openedSettingsModal = useSelector(
		(state: RootState) => state.persistedReducer.modal.openedSettingsModal
	);

	return (
		<>
			<Modal
				opened={openedSettingsModal}
				onClose={() => dispatch(setOpenSettingsModal(false))}
				title="Settings"
			>
				<Stack align="center">
					<Avatar size={"lg"} radius={"lg"} />
				</Stack>
				<SettingsModalForm />
			</Modal>
		</>
	);
};

export default SettingsModal;
