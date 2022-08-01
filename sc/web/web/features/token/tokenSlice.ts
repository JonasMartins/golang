import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface TokenState {
	value: string;
}

const initialState: TokenState = {
	value: "",
};

export const tokenSlice = createSlice({
	name: "token",
	initialState,
	reducers: {
		setToken: (state, action: PayloadAction<string>) => {
			state.value = action.payload;
		},
	},
});

export const { setToken } = tokenSlice.actions;

const tokenReducer = tokenSlice.reducer;

export default tokenReducer;
