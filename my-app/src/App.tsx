import './App.css';
import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import { ProfileHeader } from "./Profile/ProfileHeader";
import { Dashboard } from './Views/Dashboard';
import { Inbox } from './Views/Inbox';
import { EngagementReports } from './Views/EngagementReports';
import { Templates } from './Views/Templates';
import * as Constants from './Constants/constants';

export default function App() {
  return (
    <Router>
      <div className="App">
        <div className="profile-header">
          <ProfileHeader></ProfileHeader>
        </div>
        <nav className="nav-bar">
          <h1 className="title">{Constants.Title}</h1>
          <ul>
            <li><img src="./assets/icons/pie.png" /><Link className="nav-link" to="/dashboard">{Constants.Dashboard}</Link></li>
            <li><img src="./assets/icons/inbox.png" /><Link className="nav-link" to="/inbox">{Constants.Inbox}</Link></li>
            <li><img src="./assets/icons/stats.png" /><Link className="nav-link" to="/engagement-reports">{Constants.EngagementReports}</Link></li>
            <li><img src="./assets/icons/layout.png" /><Link className="nav-link" to="/templates">{Constants.Templates}</Link></li>
          </ul>
          <div className="compose-email-btn-container">
            <button className="compose-email-btn">{Constants.ComposeEmail}</button>
            <hr className="solid" />
          </div>
          <ul>
          <li><img src="./assets/icons/settings.png" /><Link className="nav-link" to="/settings">{Constants.Settings}</Link></li>
            <li><img src="./assets/icons/logout.png" /><Link className="nav-link" to="/logout">{Constants.Logout}</Link></li>
          </ul>
        </nav>
        <Route path="/" exact component={Dashboard}></Route>
        <Route path="/dashboard" exact component={Dashboard}></Route>
        <Route path="/inbox" exact component={Inbox}/>
        <Route path="/engagement-reports" exact component={EngagementReports}/>
        <Route path="/templates" exact component={Templates}/>
      </div>
    </Router>
  );
}
