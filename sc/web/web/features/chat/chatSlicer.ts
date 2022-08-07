import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType } from "@/features/types/chat";

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

		addMessage: (state, action: PayloadAction<MessageType | null>) => {
			if (action.payload && state.value) {
				console.log("Adding new message via redux");
				state.value.Messages.push(action.payload);
			}
		},
	},
});

export const { setFocusedChat, addMessage } = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
