import React, { useState } from 'react';
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
import { StatCard } from '../Dashboard/StatCard';
import './Dashboard.css';

type ChartData = {
    index: number;
    stat: number;
}

// TODO
function fetchChartData() {
     
}

function makeTestChartData(): Array<ChartData> {
    var data = [];
    for (var i = 0; i < 24; i++) {
        data.push({index: i, stat: Math.floor(Math.random() * 50)});
    }
    return data as Array<ChartData>;
}

// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    const [general, setGeneral] = useState(0);
    const [casework, setCasework] = useState(0);
    const [assigned, setAssigned] = useState(0);
    const [overdue, setOverdue] = useState(0);
    const [weeklyTotal, setWeeklyTotal] = useState(0);
    const [today, setToday] = useState(0);

    const test_data = makeTestChartData();

    const renderLineChart = (
        <div className="chart">
            <LineChart width={800} height={500} data={test_data}>
                <XAxis dataKey="index" />
                <YAxis />
                <CartesianGrid stroke="#eee" vertical={false}/>
                <Tooltip />
                <Line type="monotone" dataKey="stat" stroke="#9B51E0" />
            </LineChart>
        </div>
    );

    const renderChartCard = (
        <div className="chart-card">

        </div>
    );

    return (
        <div className="dashboard">
            <div className="stat-cards">
                <StatCard title={"General"} stat={general}></StatCard>
                <StatCard title={"Casework"} stat={casework}></StatCard>
                <StatCard title={"Assigned"} stat={assigned}></StatCard>
                <StatCard title={"Overdue"} stat={overdue}></StatCard>
            </div>
            <div className="trends">
                <h2 className="chart-title">Today's trends</h2>
                <div className="chart-stats">
                    {renderLineChart}
                    <div className="chart-stats-cards">
                        <StatCard title={"Weekly Total"} stat={weeklyTotal}></StatCard>
                        <StatCard title={"From Today"} stat={today}></StatCard>
                    </div>
                </div>
            </div>
        </div>
    );
}