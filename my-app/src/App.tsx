import './App.css';
import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import { ProfileHeader } from "./Profile/ProfileHeader";
import { Dashboard } from './Pages/Dashboard';
import { Inbox } from './Pages/Inbox';
import { EngagementReports } from './Pages/EngagementReports';
import { Templates } from './Pages/Templates';

export default function App() {
  return (
    <Router>
      <div className="App">
        <div className="profile-header">
          <ProfileHeader></ProfileHeader>
        </div>
        <nav className="nav-bar">
          <ul>
            <li><Link className="nav-link" to="/dashboard">Dashboard</Link></li>
            <li><Link className="nav-link" to="/inbox">Inbox</Link></li>
            <li><Link className="nav-link" to="/engagement-reports">Engagement Reports</Link></li>
            <li><Link className="nav-link" to="/templates">Your Templates</Link></li>
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
