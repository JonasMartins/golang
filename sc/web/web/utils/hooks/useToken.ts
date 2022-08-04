import { useState } from "react";
import { useSelector } from "react-redux";
import { RootState } from "@/app";

export const useToken = () => {
	const tokenFromStore = useSelector((state: RootState) => state.token.value);

	const [token, setTokenInternal] = useState(() => {
		return tokenFromStore;
	});

	const setToken = (newToken: string) => {
		setTokenInternal(newToken);
	};

	return [token, setToken];
};
