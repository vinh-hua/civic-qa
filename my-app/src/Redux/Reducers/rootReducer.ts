import { combineReducers } from 'redux'
import pathReducer from './pathReducer';

const rootReducer = combineReducers({
    path: pathReducer,
})

export type AppState = ReturnType<typeof rootReducer>
export default rootReducer;