import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType } from "@/features/types/chat";

export interface ChatState {
	value: ChatType | null;
}

const initialState: ChatState = {
	value: null,
};

export const chatSlice = createSlice({
	name: "chat",
	initialState,
	reducers: {
		setFocusedChat: (state, action: PayloadAction<ChatType | null>) => {
			state.value = action.payload;
		},
	},
});

export const { setFocusedChat } = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
