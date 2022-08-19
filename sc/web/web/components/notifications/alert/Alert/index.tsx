import { Alert, DefaultMantineColor } from "@mantine/core";
import React from "react";
import { IconAlertCircle, IconCheck } from "@tabler/icons";

export type AlterTypes = {
	type: "ERROR" | "SUCCESS";
};

interface GeneralMutationAlertsProps {
	type: AlterTypes;
	message: string;
	title: string;
}

const GeneralMutationAlerts: React.FC<GeneralMutationAlertsProps> = ({ type, message, title }) => {
	const handleIcon = () => {
		switch (type.type) {
			case "ERROR":
				return <IconAlertCircle size={16} />;
			case "SUCCESS":
				return <IconCheck size={16} />;
			default:
				return <IconCheck size={16} />;
		}
	};

	const handleColor = (): DefaultMantineColor => {
		switch (type.type) {
			case "ERROR":
				return "red";
			case "SUCCESS":
				return "green";
			default:
				return "green";
		}
	};

	return (
		<Alert
			icon={handleIcon()}
			title={title}
			color={handleColor()}
			withCloseButton
			closeButtonLabel="Close alert"
		>
			{message}
		</Alert>
	);
};

export default GeneralMutationAlerts;
