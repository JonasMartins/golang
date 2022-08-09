import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType } from "@/features/types/chat";

export interface ChatState {
	value: ChatType | null;
	hasAddedMessage: MessageType | null;
	searchTerm: string;
}

const initialState: ChatState = {
	value: null,
	hasAddedMessage: null,
	searchTerm: "",
};

export const chatSlice = createSlice({
	name: "chat",
	initialState,
	reducers: {
		setFocusedChat: (state, action: PayloadAction<ChatType | null>) => {
			state.value = action.payload;
		},

		addMessage: (state, action: PayloadAction<MessageType | null>) => {
			state.hasAddedMessage = action.payload;
		},

		setSearchTerm: (state, action: PayloadAction<string>) => {
			state.searchTerm = action.payload;
		},
	},
});

export const { setFocusedChat, addMessage, setSearchTerm } = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
