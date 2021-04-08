import { combineReducers } from 'redux'
import authReducer from './authReducer';
import pathReducer from './pathReducer';

const rootReducer = combineReducers({
    path: pathReducer,
    auth: authReducer,
})

export type AppState = ReturnType<typeof rootReducer>
export default rootReducer;