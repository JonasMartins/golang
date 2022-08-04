import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "@/features/counter/counterSlice";
import tokenReducer from "@/features/token/tokenSlice";
import chatReducer from "@/features/chat/chatSlicer";
import userReducer from "@/features/user/userSlice";

export const store = configureStore({
	reducer: {
		counter: counterReducer,
		token: tokenReducer,
		chat: chatReducer,
		user: userReducer,
	},
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
