import * as React from 'react';
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip, Legend } from 'recharts';
import { StatCard } from '../Dashboard/StatCard';
import './Dashboard.css';

// TODO
function fetchChartData() {
     
}

// currently using dummy data for StatCards and LineChart
export function Dashboard() {
    const test_data = [{name: 1, uv: 12},
                       {name: 2, uv: 15},
                       {name: 3, uv: 34}];

    const renderLineChart = (
        <div className="chart">
            <LineChart width={700} height={300} data={test_data}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="uv" stroke="#82ca9d" />
            </LineChart>
        </div>
    );

    return (
        <div>
            <StatCard title={"General"} stat={625}></StatCard>
            <StatCard title={"Casework"} stat={198}></StatCard>
            <StatCard title={"Assigned"} stat={190}></StatCard>
            <StatCard title={"Overdue"} stat={246}></StatCard>
            {renderLineChart}
        </div>
    );
}