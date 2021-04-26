import { useState } from 'react';
import { NavLink } from "react-router-dom";
import * as Endpoints from '../Constants/Endpoints';
import './SignupLogin.css';

export type SignupProps = {
    userLogin: Function;
}

export function Signup(props: SignupProps) {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const attemptSignup = async(e: any) => {
        e.preventDefault();
        var newUser = {email: email, password: password, passwordConfirm: confirmPassword, firstName: firstName, lastName: lastName};
        var jsonNewUser = JSON.stringify(newUser);
        const response = await fetch(Endpoints.Base + "/signup", {
            method: "POST",
            body: jsonNewUser,
            headers: new Headers({
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 400) {
            alert("Invalid email or password");
            return;
        }
        const authToken = response.headers.get("Authorization");
        props.userLogin(authToken);
    }

    function switchView() {
        setFirstName("");
        setLastName("");
        setEmail("");
        setPassword("");
        setConfirmPassword("");
    }

    return(
        <div className="signup-login-page">
            <div className="signup-login-container">
                <img className="signup-login-logo" src="./assets/icons/logo.png"></img>
                <h1 className="signup-login-title">Civic QA</h1>
                <form className="signup-login-form" autoComplete="off" onSubmit={attemptSignup}>
                    <input className="signup-login-form-input signup-login-form-text" name="fname" type="text" value={firstName} placeholder="First Name" onChange={e => setFirstName(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-form-text" name="lname" type="text" value={lastName} placeholder="Last Name" onChange={e => setLastName(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-form-text" name="email" type="text" value={email} placeholder="Email" onChange={e => setEmail(e.target.value)} required />
                    <input className="signup-login-form-input signup-login-form-text" name="password" type="password" value={password} placeholder="Password" onChange={e => setPassword(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-form-text" name="confirm-password" type="password" value={confirmPassword} placeholder="Confirm Password" onChange={e => setConfirmPassword(e.target.value)} required/>
                    <input className="signup-login-form-input signup-login-btn" type="submit" value="Sign up" />
                </form>
                <NavLink to="login" className="switch-signup-login-btn" onClick={switchView}>Already have an account? Login</NavLink>
            </div>
        </div>
    );
}