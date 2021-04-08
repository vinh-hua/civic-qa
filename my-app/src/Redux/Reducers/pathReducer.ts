import { PathActions } from "../Actions/pathActions";

type PathState = {
    path: string;
}

const initialState: PathState = {
    path: '/dashboard',
}
const pathReducer = (state: PathState = initialState, action: PathActions) => {
    switch(action.type) {
        case 'SET_PATH':
            return {
                ...state,
                path: action.payload,
            }
        default:
            return state;
    }
}

export default pathReducer;