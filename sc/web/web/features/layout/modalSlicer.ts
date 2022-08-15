import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface ModalState {
	openedSettingsModal: boolean;
}

const initialState: ModalState = {
	openedSettingsModal: false,
};

export const modalSlicer = createSlice({
	name: "modal",
	initialState,
	reducers: {
		setOpenSettingsModal: (state, action: PayloadAction<boolean>) => {
			state.openedSettingsModal = action.payload;
		},
	},
});

export const { setOpenSettingsModal } = modalSlicer.actions;

const modalReducer = modalSlicer.reducer;

export default modalReducer;
