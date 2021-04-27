import { useEffect, Dispatch } from "react";
import { NavLink, Redirect, useLocation } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from '../Redux/Reducers/rootReducer'
import { PathActions } from '../Redux/Actions/pathActions';
import { ProfileHeader } from "../Profile/ProfileHeader";
import * as Constants from '../Constants/Constants';

export type NavigationProps = {
    userLogout: Function;
}

export function Navigation(props: NavigationProps) {
    const location = useLocation();
    const { path } = useSelector((state: AppState) => state.path);
    const pathDispatch = useDispatch<Dispatch<PathActions>>();
  
    const handleSetPath = (path: string) => {
      pathDispatch({type: 'SET_PATH', payload: path})
    }
  
    useEffect(() => {
      handleSetPath(location.pathname);
    }, [location]);

    function clickLogout() {
        props.userLogout();
        <Redirect to="welcome"/>
    }

    return(
        <div>
            <div className="profile-header">
                <ProfileHeader></ProfileHeader>
            </div>
            <nav className="nav-bar">
                <h1 className="title">{Constants.Title}</h1>
                <ul>
                    <li>{path == "/dashboard" ? <img className="icon" src="./assets/icons/pie-chart-active.svg"/> : <img className="icon" src="./assets/icons/pie-chart.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/dashboard">{Constants.Dashboard}</NavLink></li>
                    <li className="dashboard-sub-li"><NavLink className="nav-link" activeClassName="active-link" to="/general">{Constants.General}</NavLink></li>
                    <li className="dashboard-sub-li"><NavLink className="nav-link" activeClassName="active-link" to="/casework">{Constants.Casework}</NavLink></li>
                    <li>{path == "/inquiries" ? <img className="icon" src="./assets/icons/inbox-active.svg"/> : <img className="icon" src="./assets/icons/inbox.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/inquiries">{Constants.Inquiries}</NavLink></li>
                    <li>{path == "/forms" ? <img className="icon" src="./assets/icons/inbox-active.svg"/> : <img className="icon" src="./assets/icons/inbox.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/forms">{Constants.Forms}</NavLink></li>
                    <li>{path =="/engagement-reports" ? <img className="icon" src="./assets/icons/stats-active.svg"/> :<img className="icon" src="./assets/icons/stats.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/engagement-reports">{Constants.EngagementReports}</NavLink></li>
                </ul>
                <div className="compose-email-btn-container">
                    <hr className="solid"/>
                </div>
                <ul>
                    <li>{path =="/settings" ? <img className="icon" src="./assets/icons/settings-active.svg"/> : <img className="icon" src="./assets/icons/settings.svg"/>}<NavLink className="nav-link" activeClassName="active-link" to="/settings">{Constants.Settings}</NavLink></li>
                    <li><img className="icon" src="./assets/icons/logout.svg"/><button className="logout-btn" onClick={clickLogout}>{Constants.Logout}</button></li>
                </ul>
            </nav>
      </div>
    );
}