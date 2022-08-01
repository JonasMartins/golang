import { useState, useEffect } from "react";
import { useToken } from "./useToken";

export type UserJwt = {
	id: string;
	name: string;
	email: string;
	exp: number;
};

export const getPayloadFromToken = (token: string): UserJwt => {
	const encodedPayload = token.split(".")[1];
	return JSON.parse(Buffer.from(encodedPayload, "base64").toString());
};

export const useUser = (): UserJwt | null => {
	const [token] = useToken();

	const [user, setUser] = useState<UserJwt | null>(() => {
		if (!token) return null;
		if (typeof token == "string") {
			return getPayloadFromToken(token);
		}
		return null;
	});

	useEffect(() => {
		if (!token) {
			setUser(null);
		} else {
			if (typeof token == "string") {
				setUser(getPayloadFromToken(token));
			}
		}
	}, [token]);

	return user;
};
