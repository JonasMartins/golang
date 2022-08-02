import { UserJwt } from "@/utils/hooks";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface UserState {
	value: UserJwt | null;
}

const initialState: UserState = {
	value: null,
};

export const userSlice = createSlice({
	name: "user",
	initialState,
	reducers: {
		setLoggedUser: (state, action: PayloadAction<UserJwt | null>) => {
			state.value = action.payload;
		},
	},
});

export const { setLoggedUser } = userSlice.actions;

const userReducer = userSlice.reducer;

export default userReducer;
