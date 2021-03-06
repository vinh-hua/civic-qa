import React, { useState } from 'react';
import './Login.css';
// import * as Endpoints from '../Endpoints/Endpoints';

export type LoginProps = {
    userLogin: Function;
}

export function Login(props: LoginProps) {
    const [onSignup, setOnSignup] = useState(false);
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const attemptLogin = async(e: any) => {
        e.preventDefault();
        var login = {email: email, password: password};
        var jsonLogin = JSON.stringify(login);
        const response = await fetch("http://localhost/v0/login", {
            method: "POST",
            body: jsonLogin,
            headers: new Headers({
                "Content-Type": "application/json"
            })
        });
        console.log(response);
        if (response.status >= 300) {
            alert("Invalid email or password");
            return;
        }
        const authToken = response.headers.get("Authorization");
        props.userLogin(authToken);
    }

    const attemptSignup = async(e: any) => {
        e.preventDefault();
        var newUser = {email: email, password: password, passwordConfirm: confirmPassword, firstName: firstName, lastName: lastName};
        var jsonNewUser = JSON.stringify(newUser);
        const response = await fetch("http://localhost/v0/signup", {
            method: "POST",
            body: jsonNewUser,
            headers: new Headers({
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            alert("Invalid email or password");
            return;
        }
        const authToken = response.headers.get("Authorization");
        props.userLogin(authToken);
    }

    function switchLoginSignup(isSignupView: boolean) {
        setFirstName("");
        setLastName("");
        setEmail("");
        setPassword("");
        setConfirmPassword("");
        setOnSignup(isSignupView);
    }

    return(
        <div className="login-page">
            <div className="login-container">
                <img src="./assets/icons/logo.png"></img>
                <h1 className="login-title">Civic QA</h1>
                {onSignup ? <div><form className="login-form" autoComplete="off" onSubmit={attemptSignup}>
                    <input className="login-form-input login-form-text" name="fname" type="text" value={firstName} placeholder="First Name" onChange={e => setFirstName(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="lname" type="text" value={lastName} placeholder="Last Name" onChange={e => setLastName(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="email" type="text" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} required />
                    <input className="login-form-input login-form-text" name="password" type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="confirm-password" type="password" value={confirmPassword} placeholder="Confirm Password" onChange={e => setConfirmPassword(e.target.value)} required/>
                    <input className="login-form-input login-btn" type="submit" value="Sign up" />
                    </form>
                    <button className="switch-login-signup-btn" onClick={() => switchLoginSignup(false)}>Already have an account? Login</button>
                </div> : 
                <div>
                    <form className="login-form" autoComplete="off" onSubmit={attemptLogin}>
                        <input className="login-form-input login-form-text" name="email" type="text" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} required/>
                        <input className="login-form-input login-form-text" name="password" type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} required/>
                        <input className="login-form-input login-btn" type="submit" value="Login" />
                    </form>
                    <button className="switch-login-signup-btn" onClick={() => switchLoginSignup(true)}>New? Create Account</button>
                </div>}
            </div>
        </div>
    );
}