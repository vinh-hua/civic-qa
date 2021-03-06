import React, { useState } from 'react';
import { StatCardRow } from '../Components/StatCardRow';
import { DashboardChartStats } from '../Dashboard/DashboardChartStats';
import * as Constants from '../Constants/Constants';
import './Dashboard.css';


// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    // stat cards data
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);
    const [topics, setTopics] = useState(0);

    let statCards = [
        {title: Constants.General, stat: general},
        {title: Constants.Casework, stat: casework},
        {title: Constants.Topics, stat: topics}
    ]

    return (
        <div className="dashboard">
            <StatCardRow cards={statCards}></StatCardRow>
            <DashboardChartStats></DashboardChartStats>
        </div>
    );
}