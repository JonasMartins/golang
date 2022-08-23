import { createSlice, createEntityAdapter } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType, MessageSubscription } from "@/features/types/chat";

export type ChatsAndUnSeenMessagesCount = {
	id: string;
	count: number;
};

export interface ChatState {
	value: ChatType | null;
	hasAddedMessage: MessageType | null;
	searchTerm: string;
	chats: ChatType[];
	chatsAndUnseenMessagesCount: ChatsAndUnSeenMessagesCount[];
}

export const chatAdapter = createEntityAdapter<ChatType>({
	selectId: chat => chat.base.id,
	sortComparer: (x, y) => (x.base.updatedAt > y.base.updatedAt ? 1 : -1),
});

const initialState: ChatState = {
	value: null,
	hasAddedMessage: null,
	searchTerm: "",
	chats: [],
	chatsAndUnseenMessagesCount: [],
};

export const chatSlice = createSlice({
	name: "chat",
	initialState,
	reducers: {
		setFocusedChat: (state, action: PayloadAction<string>) => {
			if (state.value && state.value.base.id === action.payload) return;

			let index = state.chats.findIndex(x => {
				return x.base.id === action.payload;
			});
			if (index !== -1) {
				state.value = state.chats[index];
			}
		},

		setUnseenMessagesCount: (state, action: PayloadAction<string>) => {
			let count = 0;
			for (var i = 0; i < state.chats.length; i++) {
				for (var j = 0; j < state.chats[i].Messages.length; j++) {
					if (
						state.chats[i].Messages[j].Seen &&
						!state.chats[i].Messages[j].Seen?.includes(action.payload)
					) {
						count++;
					}
				}
				if (state.chatsAndUnseenMessagesCount) {
					console.log("adding");
					state.chatsAndUnseenMessagesCount.push({
						count,
						id: action.payload,
					});
				}
				count = 0;
			}
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
		/**
		 * Every time a chat change its focus, then
		 * the most recent added message must be seted to null
		 * or will trigger the useeffect on ChatsSideBar and
		 * added a repeated message on the most recent added
		 * message's chat
		 */
		clearMessageAddedOnChangeChat: state => {
			state.hasAddedMessage = null;
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

		updadeChatsFromCommingNewMessageSubscription: (
			state,
			action: PayloadAction<MessageSubscription>
		) => {
			let index = state.chats.findIndex(x => {
				return (x.base.id = action.payload.ChatId);
			});

			if (index > -1) {
				let author = state.chats[index].Members.find(x => {
					return x.base.id === action.payload.AuthorId;
				});
				if (author && author.base.id) {
					let message: MessageType = {
						Author: {
							name: author.name,
						},
						base: action.payload.base,
						Body: action.payload.Body,
						ChatId: action.payload.ChatId,
						Seen: action.payload.Seen,
					};
					state.hasAddedMessage = message;
					state.chats[index].Messages.push(message);
				}
			}
		},
	},
	// extraReducers: buider => {
	// 	buider.addCase(PURGE, state => {
	// 		chatAdapter.removeAll()
	// 	});
	// },
});

export const {
	setFocusedChat,
	addMessage,
	setSearchTerm,
	setChats,
	clearState,
	clearMessageAddedOnChangeChat,
	updadeChatsFromCommingNewMessageSubscription,
	setUnseenMessagesCount,
} = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
