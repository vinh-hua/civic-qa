export interface ISetPathAction {
    readonly type: 'SET_PATH';
    payload: string;
}

export type PathActions =
| ISetPathAction