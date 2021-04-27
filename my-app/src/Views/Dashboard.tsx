import { useEffect, useState } from 'react';
import { StatCardRow } from '../Components/StatCardRow';
import { DashboardChartStats } from '../Dashboard/DashboardChartStats';
import * as Constants from '../Constants/Constants';
import * as Endpoints from '../Constants/Endpoints';
import './Dashboard.css';

export function Dashboard() {
    const auth = localStorage.getItem("Authorization") || "";
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);

    const getGeneralInquiries = async() => {
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
        const inquiriesGeneral = await response.json();
        setGeneral(inquiriesGeneral.length);
    }

    const getCaseworkInquiries = async() => {
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
        const inquiriesCasework = await response.json();
        setCasework(inquiriesCasework.length);
    }
    
    useEffect(() => {
        getGeneralInquiries();
        getCaseworkInquiries();
    }, []);

    let statCards = [
        {title: Constants.ActiveGeneral, stat: general},
        {title: Constants.ActiveCasework, stat: casework}
    ]

    return (
        <div className="dashboard">
            <StatCardRow spaceEven={true} cards={statCards}></StatCardRow>
            <DashboardChartStats></DashboardChartStats>
        </div>
    );
}