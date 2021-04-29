import { NavLink } from "react-router-dom";
import { TeamCard } from '../Components/TeamCard';
import * as  LandingPage from '../Constants/LandingPage';
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
                    <h1 className="landing-sub-title">Civic QA Overview</h1>
                    <p>{LandingPage.ProjectOverview1}</p>
                    <h2>Why care?</h2>
                    <p>{LandingPage.ProjectOverview2}</p>
                    <p>{LandingPage.ProjectOverview3}</p>
                </div>
                <div className="stakeholders">
                    <h1 className="landing-sub-title">Legislative Assistants</h1>
                    <p>{LandingPage.TargetStakeholders1}</p>
                    <p>{LandingPage.TargetStakeholders2}</p>
                </div>
                <div className="project-solution">
                    <h1 className="landing-sub-title">How does Civic QA help?</h1>
                    <p>{LandingPage.Solution1}</p>
                    <ol className="process-steps">
                        <li>Research</li>
                        <li>Concept Validation</li>
                        <li>Prototyping</li>
                        <li>User Testing</li>
                        <li>Feedback Implementation</li>
                    </ol>
                    <p>{LandingPage.Solution2}</p>
                    <p>{LandingPage.Solution3}</p>
                </div>
                <div className="about-team">
                    <h1 className="landing-sub-title">Meet the Team</h1>
                    <div className="team-bios">   
                        <TeamCard name="Lia Kitahata" img="./assets/images/lia.jpg" bio={LandingPage.LiaBio}></TeamCard>
                        <TeamCard name="Rafi Bayer" img="./assets/images/rafi.jpg" bio={LandingPage.RafiBio}></TeamCard>
                        <TeamCard name="Amit Galitzky" img="./assets/images/amit.jpg" bio={LandingPage.AmitBio}></TeamCard>
                        <TeamCard name="Vivian Hua" img="./assets/images/vivian.jpg" bio={LandingPage.VivianBio}></TeamCard>
                    </div>
                </div>
            </div>
        </div>
    );
}