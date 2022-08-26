import { createSlice, createEntityAdapter } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { ChatType, MessageType, MessageSubscription } from "@/features/types/chat";

export type MessagesUnreadSet = {
	messageId: string;
	seen: string[];
};

export type CacheInputsSet = {
	userId: string;
	messageId: string;
};

export interface ChatMessagesUnreadSet {
	data: Record<string, Array<MessagesUnreadSet>>;
}

export interface ChatState {
	chats: ChatType[];
	searchTerm: string;
	value: ChatType | null;
	hasAddedMessage: MessageType | null;
	recentUpdatedMessages: Array<CacheInputsSet>;
	chatsUnseenCount: ChatMessagesUnreadSet;
}

export const chatAdapter = createEntityAdapter<ChatType>({
	selectId: chat => chat.base.id,
	sortComparer: (x, y) => (x.base.updatedAt > y.base.updatedAt ? 1 : -1),
});

const initialState: ChatState = {
	chats: [],
	value: null,
	searchTerm: "",
	hasAddedMessage: null,
	recentUpdatedMessages: [],
	chatsUnseenCount: { data: {} },
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
		 * 	Recieves a userId and a message id, if this message
		 *  has not been seen, then, this method will look into
		 *  the current focused chat and will add this userId into
		 * 	message's seen array
		 * @param state
		 * @param action
		 * @returns
		 */
		updateUnseenMessage: (state, action: PayloadAction<CacheInputsSet>) => {
			for (let x of state.recentUpdatedMessages) {
				if (
					x.messageId === action.payload.messageId &&
					x.userId === action.payload.userId
				) {
					return;
				}
			}
			state.recentUpdatedMessages.push({
				messageId: action.payload.messageId,
				userId: action.payload.userId,
			});
			if (state.value && state.chatsUnseenCount.data[state.value.base.id]) {
				for (let x of state.chatsUnseenCount.data[state.value.base.id]) {
					if (x.messageId === action.payload.messageId) {
						x.seen.push(action.payload.userId);
					}
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
			/**
			 *  Set the unseen messages count when the chats comes
			 *  from first request
			 */
			let auxObj: ChatMessagesUnreadSet = { data: {} };
			// starting chatsUnseen sctructure
			for (var i = 0; i < action.payload.length; i++) {
				auxObj.data[action.payload[i].base.id] = new Array<MessagesUnreadSet>();
			}

			state.chatsUnseenCount = auxObj;
			// setting also the unseen messages struct
			for (var i = 0; i < action.payload.length; i++) {
				for (var j = 0; j < action.payload[i].Messages.length; j++) {
					if (action.payload[i].Messages[j].Seen) {
						if (
							action.payload[i].Messages[j].Seen!.length <
							action.payload[i].Members.length
						) {
							state.chatsUnseenCount.data[action.payload[i].base.id].push({
								messageId: action.payload[i].Messages[j].base.id,
								seen: action.payload[i].Messages[j].Seen || [],
							});
						}
					}
				}
			}
		},

		clearState: state => {
			state.chats = [];
			state.searchTerm = "";
			state.hasAddedMessage = null;
			state.value = null;
			state.chatsUnseenCount = { data: {} };
			state.recentUpdatedMessages = [];
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
	setChats,
	addMessage,
	clearState,
	setSearchTerm,
	setFocusedChat,
	clearMessageAddedOnChangeChat,
	updadeChatsFromCommingNewMessageSubscription,
} = chatSlice.actions;

const chatReducer = chatSlice.reducer;

export default chatReducer;
