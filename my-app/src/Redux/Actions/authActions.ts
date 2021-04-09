export interface ISetAuthAction {
    readonly type: 'SET_AUTH';
    payload: string;
}

export type AuthActions =
| ISetAuthAction