import React, { useState } from 'react';
import { StatCardRow } from '../Components/StatCardRow';
import { DashboardChartStats } from '../Dashboard/DashboardChartStats';
import { Itinerary } from '../Dashboard/Itinerary';
import * as Constants from '../Constants/constants';
import './Dashboard.css';


// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    // stat cards data
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);
    const [assigned, setAssigned] = useState(0);
    const [overdue, setOverdue] = useState(0);

    let statCards = [
        {title: Constants.General, stat: general},
        {title: Constants.Casework, stat: casework},
        {title: Constants.Assigned, stat: assigned},
        {title: Constants.Overdue, stat: overdue}
    ];

    return (
        <div className="dashboard">
            <StatCardRow cards={statCards}></StatCardRow>
            <DashboardChartStats></DashboardChartStats>
            <Itinerary></Itinerary>
        </div>
    );
}