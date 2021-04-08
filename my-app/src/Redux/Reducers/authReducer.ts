import { AuthActions } from "../Actions/authActions";

type AuthState = {
    auth: string;
}

const initialState: AuthState = {
    auth: localStorage.getItem("Authorization") || '',
}
const authReducer = (state: AuthState = initialState, action: AuthActions) => {
    switch(action.type) {
        case 'SET_AUTH':
            return {
                ...state,
                auth: action.payload,
            }
        default:
            return state;
    }
}

export default authReducer;