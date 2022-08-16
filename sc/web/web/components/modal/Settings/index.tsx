import { RootState } from "@/app";
import { setOpenSettingsModal } from "@/features/layout/modalSlicer";
import { Avatar, Modal, Stack, useMantineTheme } from "@mantine/core";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import SettingsModalForm from "@/components/form/SettingsModalForm";

interface SettingsModalProps {}

const SettingsModal: React.FC<SettingsModalProps> = ({}) => {
	const dispatch = useDispatch();
	const openedSettingsModal = useSelector(
		(state: RootState) => state.persistedReducer.modal.openedSettingsModal
	);
	const theme = useMantineTheme();

	return (
		<>
			<Modal
				opened={openedSettingsModal}
				onClose={() => dispatch(setOpenSettingsModal(false))}
				title="Settings"
				overlayColor={
					theme.colorScheme === "dark" ? theme.colors.dark[9] : theme.colors.gray[2]
				}
				overlayOpacity={0.55}
				overlayBlur={3}
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
