import React from 'react';
import StatCard from '../Dashboard/StatCard.js';

// has placeholder values for StatCards
function Dashboard() {
    return (
        <div>
            <h1>Dashboard</h1>
            <StatCard title={"General"} data={625}></StatCard>
            <StatCard title={"Casework"} data={198}></StatCard>
            <StatCard title={"Assigned"} data={190}></StatCard>
            <StatCard title={"Overdue"} data={246}></StatCard>
        </div>
    );
}

export default Dashboard;