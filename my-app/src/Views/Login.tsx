import './Login.css';

export type LoginProps = {
    login: Function;
}

export function Login(props: LoginProps) {
    function clickLoginButton() {
        // add auth verification from backend
        props.login();
        localStorage.setItem("AuthToken", "success");
    }

    return(
        <div>
            <button onClick={clickLoginButton}>Log In</button>
        </div> 
    );
}