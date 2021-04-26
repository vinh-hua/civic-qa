import { NavLink } from "react-router-dom";
import './Landing.css';

export function Landing() {
    return(
        <div className="landing-page">
            <div className="landing-logo-title-container">
                <img className="landing-logo" src="./assets/icons/logo.png"></img>
                <h1 className="landing-title">Civic QA</h1>
            </div>
            <div className="landing-btn-container">
                <NavLink className="landing-btn" to="/signup">Signup</NavLink>
                <NavLink className="landing-btn" to="login">Login</NavLink>
            </div>
            <div className="project-container">
                <div className="project-overview">
                    <h1 className="landing-sub-title">What is Civic QA?</h1>
                </div>
                <div className="stakeholders">
                    <h1 className="landing-sub-title">Legislative Assistants</h1>
                </div>
                <div className="project-solution">
                    <h1 className="landing-sub-title">How does Civic QA help?</h1>
                </div>
                <div className="about-team">
                    <h1 className="landing-sub-title">Meet the team!</h1>
                    <img className="teammate-img" src="./assets/images/rafi.jpg"></img>
                    <img className="teammate-img" src="./assets/images/amit.jpg"></img>
                    <img className="teammate-img" src="./assets/images/lia.jpg"></img>
                    <img className="teammate-img" src="./assets/images/vivian.jpg"></img>
                </div>
            </div>
        </div>
    );
}