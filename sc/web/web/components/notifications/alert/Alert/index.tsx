import { Alert, DefaultMantineColor } from "@mantine/core";
import React from "react";
import { IconAlertCircle, IconCheck } from "@tabler/icons";
import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/app";
import { closeAlert } from "@/features/alert/alertSlice";

export type AlterTypes = {
	type: "ERROR" | "SUCCESS";
};

interface GeneralMutationAlertsProps {}

const GeneralMutationAlerts: React.FC<GeneralMutationAlertsProps> = ({}) => {
	const alertOpen = useSelector((state: RootState) => state.persistedReducer.alert.open);
	const alert = useSelector((state: RootState) => state.persistedReducer.alert.alert);
	const dispatch = useDispatch();

	const handleIcon = () => {
		switch (alert!.type.type) {
			case "ERROR":
				return <IconAlertCircle size={16} />;
			case "SUCCESS":
				return <IconCheck size={16} />;
			default:
				return <IconCheck size={16} />;
		}
	};

	const handleColor = (): DefaultMantineColor => {
		switch (alert!.type.type) {
			case "ERROR":
				return "red";
			case "SUCCESS":
				return "green";
			default:
				return "green";
		}
	};

	const handleCloseAlert = () => {
		dispatch(closeAlert());
	};

	const content =
		alert && alertOpen ? (
			<Alert
				icon={handleIcon()}
				title={alert.title}
				color={handleColor()}
				withCloseButton
				closeButtonLabel="Close alert"
				onClose={handleCloseAlert}
			>
				{alert.message}
			</Alert>
		) : (
			<></>
		);

	return content;
};

export default GeneralMutationAlerts;
