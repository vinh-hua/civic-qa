import { useEffect, useState } from 'react';
import { StatCardRow } from '../Components/StatCardRow';
import { DashboardChartStats } from '../Dashboard/DashboardChartStats';
import { useSelector } from 'react-redux';
import { AppState } from '../Redux/Reducers/rootReducer'
import * as Constants from '../Constants/Constants';
import * as Endpoints from '../Constants/Endpoints';
import './Dashboard.css';

// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    const { auth } = useSelector((state: AppState) => state.auth);
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);

    const getGeneralResponses = async() => {
        const response = await fetch(Endpoints.Base + Endpoints.ResponsesActiveGeneral, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const responsesGeneral = await response.json();
        setGeneral(responsesGeneral.length);
    }

    const getCaseworkResponses = async() => {
        const response = await fetch(Endpoints.Base + Endpoints.ResponsesActiveCasework, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const responsesCasework = await response.json();
        setCasework(responsesCasework.length);
    }
    
    useEffect(() => {
        getGeneralResponses();
        getCaseworkResponses();
    }, []);

    let statCards = [
        {title: Constants.General, stat: general},
        {title: Constants.Casework, stat: casework}
    ]

    return (
        <div className="dashboard">
            <StatCardRow spaceEven={true} cards={statCards}></StatCardRow>
            <DashboardChartStats></DashboardChartStats>
        </div>
    );
}