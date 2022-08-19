import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { AlterTypes } from "@/components/notifications/alert/Alert";

export type AlertContent = {
	title: string;
	message: string;
	type: AlterTypes;
};

export interface AlertState {
	alert: AlertContent | null;
	open: boolean;
}

const initialState: AlertState = {
	alert: null,
	open: false,
};

export const alertSlice = createSlice({
	name: "alert",
	initialState,
	reducers: {
		triggerAlert: (state, action: PayloadAction<AlertContent>) => {
			(state.alert = action.payload), (state.open = true);
		},

		closeAlert: state => {
			(state.alert = null), (state.open = false);
		},
	},
});

export const { triggerAlert, closeAlert } = alertSlice.actions;

const alertReducer = alertSlice.reducer;

export default alertReducer;
