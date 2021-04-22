import './App.css';
import { useEffect, Dispatch} from "react";
import { Route, NavLink, Redirect, useLocation } from "react-router-dom";
import { ProfileHeader } from "./Profile/ProfileHeader";
import { Dashboard } from './Views/Dashboard';
import { General } from './Views/General';
import { Casework } from './Views/Casework';
import { Responses } from './Views/Responses';
import { EngagementReports } from './Views/EngagementReports';
import { Settings } from './Views/Settings';
import { Login } from './Views/Login';
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from './Redux/Reducers/rootReducer'
import { PathActions } from './Redux/Actions/pathActions';
import { AuthActions } from './Redux/Actions/authActions';
import * as Constants from './Constants/Constants';
import * as Endpoints from './Constants/Endpoints';

export default function App() {
  const location = useLocation();
  const { auth } = useSelector((state: AppState) => state.auth);
  const { path } = useSelector((state: AppState) => state.path);
  const authDispatch = useDispatch<Dispatch<AuthActions>>();
  const pathDispatch = useDispatch<Dispatch<PathActions>>();

  const handleSetAuth = (auth: string) => {
    authDispatch({type: 'SET_AUTH', payload: auth})
  }

  const handleSetPath = (path: string) => {
    pathDispatch({type: 'SET_PATH', payload: path})
  }

  useEffect(() => {
    handleSetPath(location.pathname);
  }, [location]);

  const userLogout = async(e: any) => {
    e.preventDefault();
    const response = await fetch(Endpoints.Base + "/logout", {
      method: "POST",
      headers: new Headers({
          "Authorization": auth
      })
    });
    if (response.status >= 300) {
      console.log("Failed to logout");
    }
    handleSetAuth('');
  }

  function userLogin(authToken: string) {
    localStorage.setItem("Authorization", authToken);
    handleSetAuth(authToken);
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
            <Route path="/dashboard" component={Dashboard}/>
            <Route path="/general" component={General}/>
            <Route path="/casework" component={Casework}/>
            <Route path="/responses" component={Responses}/>
            <Route path="/engagement-reports" component={EngagementReports}/>
            <Route path="/settings" component={Settings}/>
          </div> : null}
        </div>
  );
}
