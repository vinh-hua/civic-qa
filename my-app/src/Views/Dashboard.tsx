import React, { useState } from 'react';
import { AreaChart, Area, LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
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
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);
    const [assigned, setAssigned] = useState(0);
    const [overdue, setOverdue] = useState(0);
    const [total, setTotal] = useState(0);
    const [today, setToday] = useState(0);

    const test_data = makeTestChartData();

    const renderAreaChart = (
        <div>
            <AreaChart width={800} height={500} data={test_data}>
            <defs>
                <linearGradient id="purpleGradient" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#9B51E0" stopOpacity={0.8}/>
                    <stop offset="95%" stopColor="#9B51E0" stopOpacity={0}/>
                </linearGradient>
            </defs>
            <XAxis dataKey="name" />
            <YAxis />
            <CartesianGrid stroke="#eee" vertical={false} />
            <Tooltip />
            <Area type="monotone" dataKey="count" stroke="#9B51E0" fillOpacity={1} fill="url(#purpleGradient)" />
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
                <h2 className="chart-title">{Constants.ChartTitle}</h2>
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