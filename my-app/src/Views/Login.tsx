import React, { useState } from 'react';
import './Login.css';
// import * as Endpoints from '../Endpoints/Endpoints';

export type LoginProps = {
    login: Function;
}

async function userSignup() {
    var newUser = {email: "test@test.com", password: "abcdefgh", passwordConfirm: "abcdefgh", firstName: "vivian", lastName: "hua"};
    var jsonNewUser = JSON.stringify(newUser);
    const response = await fetch("http://localhost/v0/signup", {
        method: "POST",
        body: jsonNewUser,
        headers: new Headers({
            "Content-Type": "application/json"
        })
    })
    console.log(response);
}

async function userLogin() {

}

export function Login(props: LoginProps) {
    const [onSignup, setOnSignup] = useState(false);
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    function userLogin() {
        // add auth verification from backend
        props.login();
        localStorage.setItem("user", "success");
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
                {onSignup ? <div><form className="login-form" autoComplete="off" onSubmit={userSignup}>
                    <input className="login-form-input login-form-text" name="fname" type="text" value={firstName} placeholder="First Name" onChange={e => setFirstName(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="lname" type="text" value={lastName} placeholder="Last Name" onChange={e => setLastName(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="email" type="text" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} required />
                    <input className="login-form-input login-form-text" name="password" type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} required/>
                    <input className="login-form-input login-form-text" name="password" type="password" value={password} placeholder="Confirm Password" onChange={e => setConfirmPassword(e.target.value)} required/>
                    <input className="login-form-input login-btn" type="submit" value="Sign up" />
                    </form>
                    <button className="switch-login-signup-btn" onClick={() => switchLoginSignup(false)}>Already have an account? Login</button>
                </div> : 
                <div>
                    <form className="login-form" autoComplete="off" onSubmit={userLogin}>
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