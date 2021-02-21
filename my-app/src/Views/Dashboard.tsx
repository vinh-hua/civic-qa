import React, { useState } from 'react';
import { DropdownMenu } from '../Dashboard/DropdownMenu';
import { StatCard } from '../Dashboard/StatCard';
import { ChartData, DashboardChart } from '../Dashboard/DashboardChart';
import * as Constants from '../Constants/constants';
import './Dashboard.css';

function randomChartData(): Array<ChartData> {
    var data = [];
    for (var i = 0; i < 24; i++) {
        data.push({index: i, count: Math.floor(Math.random() * 50)});
    }
    return data as Array<ChartData>;
}

// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    // stat cards data
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);
    const [assigned, setAssigned] = useState(0);
    const [overdue, setOverdue] = useState(0);
    const [total, setTotal] = useState(0);
    const [today, setToday] = useState(0);
    const [chartView, setChartView] = useState(Constants.AllEmails);

    // make randomize chart data
    const test_data = randomChartData();

    return (
        <div className="dashboard">
            <div className="stat-cards">
                <StatCard title={Constants.General} stat={general}></StatCard>
                <StatCard title={Constants.Casework} stat={casework}></StatCard>
                <StatCard title={Constants.Assigned} stat={assigned}></StatCard>
                <StatCard title={Constants.Overdue} stat={overdue}></StatCard>
            </div>
            <div className="trends">
                <div className="chart-heading">
                    <h2 className="chart-title">{Constants.ChartTitle}</h2>
                    <div className="dropdown-menu">
                        <DropdownMenu chartView={chartView} setChartView={setChartView}></DropdownMenu>
                    </div>
                </div>
                <div className="chart-stats">
                    <DashboardChart data={test_data}></DashboardChart>
                    <div className="chart-stats-cards">
                        <StatCard title={Constants.Total} stat={total}></StatCard>
                        <StatCard title={Constants.Today} stat={today}></StatCard>
                    </div>
                </div>
            </div>
        </div>
    );
}