import './App.css';
import React, { useEffect, useState} from "react";
import { Route, NavLink, Redirect, useLocation } from "react-router-dom";
import { ProfileHeader } from "./Profile/ProfileHeader";
import { Dashboard } from './Views/Dashboard';
import { General } from './Views/General';
import { Casework } from './Views/Casework';
import { Responses } from './Views/Responses';
import { EngagementReports } from './Views/EngagementReports';
import { Templates } from './Views/Templates';
import { Settings } from './Views/Settings';
import { Login } from './Views/Login';
import * as Constants from './Constants/Constants';

export default function App() {
  const authToken = localStorage.getItem("Authorization");
  const [auth, setAuth] = useState((authToken != "") && (authToken != null));
  const [path, setPath] = useState("/dashboard");
  const location = useLocation();

  useEffect(() => {
    setPath(location.pathname);
  }, [location]);

  const userLogout = async(e: any) => {
    e.preventDefault();
    var authToken = localStorage.getItem("Authorization") || "";
    const response = await fetch("http://localhost/v0/logout", {
      method: "POST",
      headers: new Headers({
          "Authorization": authToken
      })
    });
    if (response.status >= 300) {
      console.log("Failed to logout");
    }
    localStorage.removeItem("Authorization");
    setAuth(false);
  }

  function userLogin(authToken: string) {
    localStorage.setItem("Authorization", authToken);
    setAuth(true);
  }

  return (
      <div className="App">
        <Route path="/login" component={() => <Login userLogin={userLogin}/>}></Route>
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
                  <li>{path == "/dashboard" ? <img className="icon" src="./assets/icons/pie-chart-active.svg"/> : <img className="icon" src="./assets/icons/pie-chart.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/dashboard">{Constants.Dashboard}</NavLink></li>
                  <li className="dashboard-sub-li"><NavLink className="nav-link" activeClassName="active-link" to="/general">{Constants.GeneralInquiries}</NavLink></li>
                  <li className="dashboard-sub-li"><NavLink className="nav-link" activeClassName="active-link" to="/casework">{Constants.Casework}</NavLink></li>
                  <li>{path =="/responses" ? <img className="icon" src="./assets/icons/inbox-active.svg"/> : <img className="icon" src="./assets/icons/inbox.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/responses">{Constants.Responses}</NavLink></li>
                  <li>{path =="/engagement-reports" ? <img className="icon" src="./assets/icons/stats-active.svg"/> :<img className="icon" src="./assets/icons/stats.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/engagement-reports">{Constants.EngagementReports}</NavLink></li>
                </ul>
                <div className="compose-email-btn-container">
                  <hr className="solid"/>
                </div>
                <ul>
                <li>{path =="/settings" ? <img className="icon" src="./assets/icons/settings-active.svg"/> : <img className="icon" src="./assets/icons/settings.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/settings">{Constants.Settings}</NavLink></li>
                  <li><img className="icon" src="./assets/icons/logout.svg"/><button className="logout-btn" onClick={userLogout}>{Constants.Logout}</button></li>
                </ul>
              </nav>
            </div>
            <Route path="/dashboard" component={Dashboard}></Route>
            <Route path="/general" component={General}></Route>
            <Route path="/casework" component={Casework}></Route>
            <Route path="/responses" component={Responses}/>
            <Route path="/engagement-reports" component={EngagementReports}/>
            <Route path="/settings" component={Settings}/>
          </div> : null}
        </div>
  );
}
