import { combineReducers, configureStore } from "@reduxjs/toolkit";
import counterReducer from "@/features/counter/counterSlice";
import tokenReducer from "@/features/token/tokenSlice";
import chatReducer from "@/features/chat/chatSlicer";
import userReducer from "@/features/user/userSlice";
import modalReducer from "@/features/layout/modalSlicer";
import storage from "redux-persist/lib/storage";
import {
	persistStore,
	persistReducer,
	FLUSH,
	REHYDRATE,
	PAUSE,
	PERSIST,
	PURGE,
	REGISTER,
} from "redux-persist";

const rootReducer = combineReducers({
	counter: counterReducer,
	token: tokenReducer,
	chat: chatReducer,
	user: userReducer,
	modal: modalReducer,
});

const persistConfig = {
	key: "root",
	version: 1,
	storage,
	blacklist: ["navigation", "counter", "layout"],
};

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
	reducer: {
		persistedReducer,
	},
	middleware: getDefaultMiddleware =>
		getDefaultMiddleware({
			serializableCheck: {
				ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
			},
			ignoredActionPaths: ["meta.arg", "payload.timestamp"],
			ignoredPaths: ["items.dates"],
		}),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

export const persistor = persistStore(store);
