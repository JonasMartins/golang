import { useState, useEffect } from "react";
import { useToken } from "./useToken";

export type User = {
	Id: string;
	Name: string;
	Email: string;
};

export const useUser = (): User | null => {
	const [token] = useToken();

	const getPayloadFromToken = (token: string): User => {
		const encodedPayload = token.split(".")[1];
		return JSON.parse(Buffer.from(encodedPayload, "base64").toString());
	};

	const [user, setUser] = useState<User | null>(() => {
		if (!token) return null;
		return getPayloadFromToken(token);
	});

	useEffect(() => {
		if (!token) {
			setUser(null);
		} else {
			setUser(getPayloadFromToken(token));
		}
	}, [token]);

	return user;
};
