import { useState, useEffect } from "react";
import { useToken } from "./useToken";

export type User = {
	Id: string;
	Name: string;
	Email: string;
	exp: number;
};

export const getPayloadFromToken = (token: string): User => {
	const encodedPayload = token.split(".")[1];
	return JSON.parse(Buffer.from(encodedPayload, "base64").toString());
};

export const useUser = (): User | null => {
	const [token] = useToken();

	const [user, setUser] = useState<User | null>(() => {
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
