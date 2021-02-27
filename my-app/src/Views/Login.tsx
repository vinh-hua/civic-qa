import './Login.css';

export type LoginProps = {
    login: Function;
}

export function Login(props: LoginProps) {
    function clickLogin() {
        // add auth verification from backend
        props.login();
        localStorage.setItem("user", "success");
    }

    return(
        <div className="login-page">
            <div className="login-container">
                <img src="./assets/icons/logo.png"></img>
                <h1 className="login-title">Civic QA</h1>
                <form className="login-form" onSubmit={clickLogin}>
                    <input className="login-form-input" type="text" placeholder="Email" required/>
                    <input className="login-form-input" type="password" placeholder="Password" required/>
                    <input className="login-form-input login-btn" type="submit" value="Login" />
                </form>
            </div>
        </div> 
    );
}