import { Alert } from "@mantine/core";
import React from "react";

export type AlterTypes = {
	type: "ERROR" | "SUCCESS";
};

interface GeneralMutationAlertsProps {
	type: AlterTypes;
	message: string;
}

const GeneralMutationAlerts: React.FC<GeneralMutationAlertsProps> = ({ type, message }) => {
	return <Alert></Alert>;
};

export default GeneralMutationAlerts;
