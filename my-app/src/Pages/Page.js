import React from 'react';
import Header from '../Navigation/Header.js';
import Stat from '../Dashboard/Stat.js';
import DashboardChart from '../Dashboard/DashboardChart.js';

class Page extends React.Component {
    render() {
        return (
            <div>
                <Header></Header>
                <h1>Dashboard</h1>
                <Stat></Stat>
                <DashboardChart></DashboardChart>
            </div>
        );
    }
}

export default Page;