import './App.css';
import { useState} from "react";

import { Route, Redirect } from "react-router-dom";
import { Navigation } from './Components/Navigation';
import { Dashboard } from './Views/Dashboard';
import { General } from './Views/General';
import { Casework } from './Views/Casework';
import {Inquiries } from './Views/Inquiries';
import { Forms } from './Views/Forms';
import { EngagementReports } from './Views/EngagementReports';
import { Settings } from './Views/Settings';
import { Landing } from './Views/Landing';
import { Signup } from './Views/Signup';
import { Login } from './Views/Login'
import * as Endpoints from './Constants/Endpoints';

export default function App() {
  const authToken = localStorage.getItem("Authorization");
  const [auth, setAuth] = useState(authToken || "");

  async function userLogout() {
    const response = await fetch(Endpoints.Base + "/logout", {
      method: "POST",
      headers: new Headers({
          "Authorization": auth
      })
    });
    if (response.status >= 400) {
      console.log("Failed to logout");
    }
    localStorage.removeItem("Authorization");
    setAuth("");
  }

  function userLogin(authToken: string) {
    localStorage.setItem("Authorization", authToken);
    setAuth(authToken);
    console.log(auth);
  }

  return (
      <div className="App">
        <Route exact path="/" component={Landing}></Route>
        <Route exact path="/signup" component={() => <Signup userLogin={userLogin}/>}></Route>
        <Route exact path="/login" component={() => <Login userLogin={userLogin}/>}></Route>
        {auth ? <Redirect to="/dashboard"/> : <Redirect to="/"/>}
        {auth ?
          <div>
            <Navigation userLogout={userLogout}></Navigation>
            <div className="view">
              <Route exact path="/dashboard" component={Dashboard}/>
              <Route exact path="/general" component={General}/>
              <Route exact path="/casework" component={Casework}/>
              <Route exact path="/inquiries" component={Inquiries}/>
              <Route exact path="/forms" component={Forms}/>
              <Route exact path="/engagement-reports" component={EngagementReports}/>
              <Route exact path="/settings" component={Settings}/>
            </div>
          </div> : null}
        </div>
  );
}
