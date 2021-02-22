import React, { useState } from 'react';
import { StatCard } from '../Components/StatCard';
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

    return (
        <div className="dashboard">
            <div className="stat-cards">
                <StatCard title={Constants.General} stat={general}></StatCard>
                <StatCard title={Constants.Casework} stat={casework}></StatCard>
                <StatCard title={Constants.Assigned} stat={assigned}></StatCard>
                <StatCard title={Constants.Overdue} stat={overdue}></StatCard>
            </div>
            <DashboardChartStats></DashboardChartStats>
            <Itinerary></Itinerary>
        </div>
    );
}