import React, { useEffect, useState } from 'react';
import { DropdownMenu } from '../Components/DropdownMenu';
import { ChartData, DashboardChart } from '../Dashboard/DashboardChart';
import { SubDashboardData } from '../Components/SubDashboard';
import { StatCard } from '../Components/StatCard';
import * as Constants from '../Constants/Constants';
import * as Endpoints from '../Constants/Endpoints';

function randomChartData(): Array<ChartData> {
    var data = [];
    for (var i = 0; i < 24; i++) {
        data.push({index: i, count: Math.floor(Math.random() * 50)});
    }
    return data as Array<ChartData>;
}

export function DashboardChartStats() {
    const [total, setTotal] = useState(0);
    const [today, setToday] = useState(0);
    const [chartView, setChartView] = useState(Constants.Responses);

    // make randomize chart data
    const test_data = randomChartData();

    async function GetStats() {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const forms = await response.json();
        setTotal(forms.length);
    }

    useEffect(() => {
        GetStats();
    }, []);

    return(
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
    );
}