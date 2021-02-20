import React, { useState } from 'react';
import { AreaChart, Area, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
import { Dropdown } from '../Components/Dropdown';
import { StatCard } from '../Dashboard/StatCard';
import * as Constants from '../Constants/constants';
import './Dashboard.css';

type ChartData = {
    index: number;
    count: number;
}

// TODO
function fetchChartData() {
     
}

function makeTestChartData(): Array<ChartData> {
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
    // dropdown menu state
    const [showMenu, toggleMenu] = useState(false);

    const test_data = makeTestChartData();

    const renderAreaChart = (
        <div>
            <AreaChart width={800} height={500} data={test_data}>
            <defs>
                <linearGradient id="purpleGradient" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="0%" stopColor={Constants.Purple} stopOpacity={0.5}/>
                    <stop offset="100%" stopColor={Constants.Purple} stopOpacity={0}/>
                </linearGradient>
            </defs>
            <XAxis dataKey="name" />
            <YAxis />
            <CartesianGrid stroke="#eee" vertical={false} />
            <Tooltip />
            <Area type="monotone" dataKey="count" stroke={Constants.Purple} fillOpacity={1} fill="url(#purpleGradient)" />
            </AreaChart>
        </div>
    );

    return (
        <div className="dashboard">
            <div className="stat-cards">
                <StatCard title={Constants.General} stat={general}></StatCard>
                <StatCard title={Constants.Casework} stat={casework}></StatCard>
                <StatCard title={Constants.Assigned} stat={assigned}></StatCard>
                <StatCard title={Constants.Overdue} stat={overdue}></StatCard>
            </div>
            <div className="trends">
                <div className="chart-title">
                    <h2>{Constants.ChartTitle}</h2>
                </div>
                <div className="chart-stats">
                    {renderAreaChart}
                    <div className="chart-stats-cards">
                        <StatCard title={Constants.Total} stat={total}></StatCard>
                        <StatCard title={Constants.Today} stat={today}></StatCard>
                    </div>
                </div>
            </div>
        </div>
    );
}