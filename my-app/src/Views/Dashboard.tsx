import React, { useEffect, useState } from 'react';
import { StatCardRow } from '../Components/StatCardRow';
import { DashboardChartStats } from '../Dashboard/DashboardChartStats';
import * as Constants from '../Constants/Constants';
import * as Endpoints from '../Constants/Endpoints';
import './Dashboard.css';

// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    // stat cards data
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);

    const getGeneralResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + Endpoints.ResponsesActiveGeneral, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
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
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + Endpoints.ResponsesActiveCasework, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
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
            <StatCardRow cards={statCards}></StatCardRow>
            <DashboardChartStats></DashboardChartStats>
        </div>
    );
}