import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType } from "@/features/types/chat";

export type PushNewMessage = {
	messageCount: number;
	chatId: string;
};

export interface ChatState {
	value: ChatType | null;
	hasAddedMessage: PushNewMessage;
}

const initialState: ChatState = {
	value: null,
	hasAddedMessage: { messageCount: 0, chatId: "" },
};

export const chatSlice = createSlice({
	name: "chat",
	initialState,
	reducers: {
		setFocusedChat: (state, action: PayloadAction<ChatType | null>) => {
			state.value = action.payload;
			state.hasAddedMessage = { messageCount: 0, chatId: action.payload?.base.id || "" };
		},

		addMessage: (state, action: PayloadAction<MessageType | null>) => {
			if (action.payload && state.value) {
				state.value.Messages.push(action.payload);
				state.hasAddedMessage = {
					messageCount: state.hasAddedMessage.messageCount + 1,
					chatId: state.value.base.id,
				};
			}
		},
	},
});

export const { setFocusedChat, addMessage } = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
