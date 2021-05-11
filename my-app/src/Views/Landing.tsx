import { NavLink } from "react-router-dom";
import { TeamCard } from '../Components/TeamCard';
import * as  LandingPage from '../Constants/LandingPage';
import './Landing.css';

export function Landing() {
    return(
        <div className="landing-page">
            <div className="landing-page-bar">
                <div className="landing-logo-title-container">
                    <img className="landing-logo" src="./assets/icons/logo.png"></img>
                    <h1 className="landing-title">Civic QA</h1>
                </div>
                <div className="landing-btn-container">
                    <NavLink className="landing-btn" to="/signup">Signup</NavLink>
                    <NavLink className="landing-btn" to="login">Login</NavLink>
                </div>
            </div>
            <div className="project-container">
                <div className="landing-sub-section project-overview">
                    <h1 className="landing-sub-title">Problem Overview</h1>
                    <p>{LandingPage.ProblemOverview1}</p>
                    <p>{LandingPage.ProblemOverview2}</p>
                    <p>{LandingPage.ProblemOverview3}</p>
                </div>
                <div className="landing-sub-section stakeholders">
                    <h1 className="landing-sub-title">Legislative Assistants</h1>
                    <p>{LandingPage.TargetStakeholders1}</p>
                    <p>{LandingPage.TargetStakeholders2}</p>
                </div>
                <div className="landing-sub-section project-solution">
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
                <div className="landing-sub-section presentation-video">
                    <h1 className="landing-sub-title">Presentation Video</h1>
                    <iframe height="500" width="800" src="https://www.youtube.com/embed/COTimhg6Cs8"></iframe>
                </div>
                <div className="landing-sub-section about-team">
                    <h1 className="landing-sub-title">Meet the Team</h1>
                    <div className="team-bios">   
                        <TeamCard name="Lia Kitahata" img="./assets/images/lia.jpg" bio={LandingPage.LiaBio} linkedin="https://www.linkedin.com/in/liakitahata/"></TeamCard>
                        <TeamCard name="Rafi Bayer" img="./assets/images/rafi.jpg" bio={LandingPage.RafiBio} linkedin="https://www.linkedin.com/in/rafael-bayer/"></TeamCard>
                        <TeamCard name="Amit Galitzky" img="./assets/images/amit.jpg" bio={LandingPage.AmitBio} linkedin="https://www.linkedin.com/in/amit-galitzky-844272171/"></TeamCard>
                        <TeamCard name="Vivian Hua" img="./assets/images/vivian.jpg" bio={LandingPage.VivianBio} linkedin="https://www.linkedin.com/in/viviancarolinehua"></TeamCard>
                    </div>
                </div>
                <div className="landing-sub-section project-status">
                    <h1 className="landing-sub-title">Project Status</h1>
                    <p>{LandingPage.ProjectStatus} <a href="https://github.com/Team-RAVL/civic-qa">here</a>.</p>
                </div>
            </div>
        </div>
    );
}