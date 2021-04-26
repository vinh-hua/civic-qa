import { useState } from 'react';
import { NavLink } from "react-router-dom";
import * as Endpoints from '../Constants/Endpoints';
import './SignupLogin.css';

export type LoginProps = {
    userLogin: Function;
}

export function Login(props: LoginProps) {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const attemptLogin = async(e: any) => {
        e.preventDefault();
        var login = {email: email, password: password};
        var jsonLogin = JSON.stringify(login);
        const response = await fetch(Endpoints.Base + "/login", {
            method: "POST",
            body: jsonLogin,
            headers: new Headers({
                "Content-Type": "application/json"
            })
        });
        console.log(response);
        if (response.status >= 400) {
            alert("Invalid email or password");
            return;
        }
        const authToken = response.headers.get("Authorization");
        props.userLogin(authToken);
    }

    function switchView() {
        setEmail("");
        setPassword("");
    }

    return(
        <div className="signup-login-page">
            <div className="signup-login-container">
                <img className="signup-login-logo" src="./assets/icons/logo.png"></img>
                <h1 className="signup-login-title">Civic QA</h1>
                <form className="signup-login-form" autoComplete="off" onSubmit={attemptLogin}>
                    <input className="signup-login-form-input signup-login-form-text" name="email" type="text" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-form-text" name="password" type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-btn" type="submit" value="Login" />
                </form>
                <NavLink to="/signup" className="switch-signup-login-btn" onClick={switchView}>New? Create Account</NavLink>
            </div>
        </div>
    );
}