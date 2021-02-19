import * as React from 'react';
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
    const test_data = makeTestChartData();

    const renderLineChart = (
        <div className="chart">
            <LineChart width={700} height={300} data={test_data}>
                <XAxis dataKey="index" />
                <YAxis />
                <CartesianGrid stroke="#eee" vertical={false}/>
                <Tooltip />
                <Line type="monotone" dataKey="stat" stroke="#3751FF" />
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
                <StatCard title={"General"} stat={625}></StatCard>
                <StatCard title={"Casework"} stat={198}></StatCard>
                <StatCard title={"Assigned"} stat={190}></StatCard>
                <StatCard title={"Overdue"} stat={246}></StatCard>
            </div>
            <div className="chart-stats">
                <h2 className="chart-title">Today's trends</h2>
                {renderLineChart}
            </div>
        </div>
    );
}