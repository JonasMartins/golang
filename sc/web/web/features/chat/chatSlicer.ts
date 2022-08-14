import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType } from "@/features/types/chat";

export interface ChatState {
	value: ChatType | null;
	hasAddedMessage: MessageType | null;
	searchTerm: string;
	chats: ChatType[];
}

const initialState: ChatState = {
	value: null,
	hasAddedMessage: null,
	searchTerm: "",
	chats: [],
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
			if (action.payload) {
				let index = state.chats.findIndex(x => {
					return x.base.id === action.payload?.ChatId;
				});

				if (index !== -1 && state.chats.length >= index) {
					state.chats[index].Messages.push(action.payload);
				}
			}
		},

		setSearchTerm: (state, action: PayloadAction<string>) => {
			state.searchTerm = action.payload;
		},

		setChats: (state, action: PayloadAction<ChatType[]>) => {
			state.chats = action.payload;
		},

		clearState: state => {
			state.chats = [];
			state.searchTerm = "";
			state.hasAddedMessage = null;
			state.value = null;
		},
	},
});

export const { setFocusedChat, addMessage, setSearchTerm, setChats, clearState } =
	chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
