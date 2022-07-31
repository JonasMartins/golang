import cookieCutter from "cookie-cutter";
import { useState } from "react";
import { COOKIE_NAME } from "@/utils/consts";

export const useToken = () => {
	const [token, setTokenInternal] = useState(() => {
		return cookieCutter.get(COOKIE_NAME);
	});

	const setToken = (newToken: string) => {
		cookieCutter.set(COOKIE_NAME, newToken);
		setTokenInternal(newToken);
	};

	return [token, setToken];
};
