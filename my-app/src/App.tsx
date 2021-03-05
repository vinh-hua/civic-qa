import './App.css';
import React, { useState} from "react";
import { BrowserRouter as Router, Route, Link, Redirect } from "react-router-dom";
import { ProfileHeader } from "./Profile/ProfileHeader";
import { Dashboard } from './Views/Dashboard';
import { General } from './Views/General';
import { Casework } from './Views/Casework';
import { Responses } from './Views/Responses';
import { EngagementReports } from './Views/EngagementReports';
import { Templates } from './Views/Templates';
import { Settings } from './Views/Settings';
import { Login } from './Views/Login';
import * as Constants from './Constants/constants';

export default function App() {
  const authToken = localStorage.getItem("user");
  const [auth, setAuth] = useState((authToken != "") && (authToken != null));

  function UserLogout() {
      setAuth(false);
      localStorage.setItem("user", "");
      return(<Redirect to="/login"></Redirect>);
  }

  function UserLogin() {
    setAuth(true);
    localStorage.setItem("user", "success");
    return(<Redirect to="/dashboard"></Redirect>);
  }

  return (
    <Router>
      <div className="App">
        <Route path="/login" component={() => <Login login={UserLogin}/>}></Route>
        {auth ? <Redirect to="/dashboard"/> : <Redirect to="/login"/>}
        {auth ?
          <div>
            <div>
              <div className="profile-header">
                <ProfileHeader></ProfileHeader>
              </div>
              <nav className="nav-bar">
                <h1 className="title">{Constants.Title}</h1>
                <ul>
                  <li><img src="./assets/icons/pie.png"/><Link className="nav-link" to="/dashboard">{Constants.Dashboard}</Link></li>
                  <li className="dashboard-sub-li"><Link className="nav-link" to="/general">{Constants.GeneralInquiries}</Link></li>
                  <li className="dashboard-sub-li"><Link className="nav-link" to="/casework">{Constants.Casework}</Link></li>
                  <li><img src="./assets/icons/inbox.png"/><Link className="nav-link" to="/responses">{Constants.Responses}</Link></li>
                  <li><img src="./assets/icons/stats.png"/><Link className="nav-link" to="/engagement-reports">{Constants.EngagementReports}</Link></li>
                  <li><img src="./assets/icons/layout.png"/><Link className="nav-link" to="/templates">{Constants.Templates}</Link></li>
                </ul>
                <div className="compose-email-btn-container">
                  <hr className="solid"/>
                </div>
                <ul>
                <li><img src="./assets/icons/settings.png"/><Link className="nav-link" to="/settings">{Constants.Settings}</Link></li>
                  <li><img src="./assets/icons/logout.png"/><button className="logout-btn" onClick={UserLogout}>{Constants.Logout}</button></li>
                </ul>
              </nav>
            </div>
            <Route path="/dashboard" component={Dashboard}></Route>
            <Route path="/general" component={General}></Route>
            <Route path="/casework" component={Casework}></Route>
            <Route path="/responses" component={Responses}/>
            <Route path="/engagement-reports" component={EngagementReports}/>
            <Route path="/templates" component={Templates}/>
            <Route path="/settings" component={Settings}/>
          </div> : null}
        </div>
      </Router>
  );
}
