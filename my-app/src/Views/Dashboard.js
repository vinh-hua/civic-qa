import React from 'react';
import { ResponsiveContainer, LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip, Legend } from 'recharts';
import StatCard from '../Dashboard/StatCard.js';
import './Dashboard.css';

// currently using dummy data for StatCards and LineChart
function Dashboard() {
    const test_data = [{name: 1, uv: 12},
                       {name: 2, uv: 15},
                       {name: 3, uv: 34}];
    const renderLineChart = (
        <div class="chart">
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
            <StatCard title={"General"} data={625}></StatCard>
            <StatCard title={"Casework"} data={198}></StatCard>
            <StatCard title={"Assigned"} data={190}></StatCard>
            <StatCard title={"Overdue"} data={246}></StatCard>
            {renderLineChart}
        </div>
    );
}

export default Dashboard;