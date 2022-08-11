import {configureStore, combineReducers} from '@reduxjs/toolkit';
import profileReducer from './slices/profileSlice';
import uiReducer from './slices/uiSlice';
import storage from 'redux-persist/lib/storage';
import {
  persistReducer,
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
  persistStore,
} from 'redux-persist';

const rootReducer = combineReducers({
  ui: uiReducer,
  profile: profileReducer,
});

const slicesToPersist = ['ui'];

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof rootReducer>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;

const persistConfig = {
  key: 'root',
  version: 1,
  storage,
  whitelist: slicesToPersist,
};

const persistedReducer = persistReducer(persistConfig, rootReducer);

const store = configureStore({
  reducer: persistedReducer,
  devTools: process.env.NODE_ENV !== 'production',
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

const defaultExports = () => {
  const persistor = persistStore(store);
  return {store, persistor};
};

export default defaultExports;
