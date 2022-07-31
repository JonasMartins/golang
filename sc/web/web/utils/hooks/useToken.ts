import { useState } from "react";

export const useToken = () => {
	const [token, setTokenInternal] = useState(() => {
		return "";
	});

	const setToken = (newToken: string) => {
		setTokenInternal(newToken);
	};

	return [token, setToken];
};
