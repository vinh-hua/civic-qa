import './App.css';
import { useState} from "react";
import { Route, NavLink, Redirect } from "react-router-dom";
import { Navigation } from './Components/Navigation';
import { Dashboard } from './Views/Dashboard';
import { General } from './Views/General';
import { Casework } from './Views/Casework';
import {Inquiries } from './Views/Inquiries';
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
        {auth ? <Redirect to="/dashboard"/> : <Redirect to="/welcome"/>}
        <Route path="/welcome" component={Landing}></Route>
        <Route path="/signup" component={() => <Signup userLogin={userLogin}/>}></Route>
        <Route path="/login" component={() => <Login userLogin={userLogin}/>}></Route>
        {auth ?
          <div>
            <Navigation userLogout={userLogout}></Navigation>
            <div className="view">
              <Route path="/signup" exact component={Signup}></Route>
              <Route path="/login" exact component={Login}></Route>
              <Route path="/dashboard" exact component={Dashboard}/>
              <Route path="/general" exact component={General}/>
              <Route path="/casework" exact component={Casework}/>
              <Route path="/inquiries" exact component={Inquiries}/>
              <Route path="/engagement-reports" exact component={EngagementReports}/>
              <Route path="/settings" exact component={Settings}/>
            </div>
          </div> : 
          <div>
            <NavLink to="/signup">Signup</NavLink>
            <NavLink to="login">Login</NavLink>
          </div>}
        </div>
  );
}
