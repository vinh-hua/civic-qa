import React, { useEffect, useState } from 'react';
import { DropdownMenu } from '../Components/DropdownMenu';
import { ChartData, DashboardChart } from '../Dashboard/DashboardChart';
import { StatCard } from '../Components/StatCard';
import * as Constants from '../Constants/Constants';
import * as Endpoints from '../Constants/Endpoints';

export function DashboardChartStats() {
    const [total, setTotal] = useState(0);
    const [today, setToday] = useState(0);
    const [todayTrends, setTodayTrends] = useState<ChartData[]>([]);
    const [chartView, setChartView] = useState(Constants.Responses);

    async function getTotal() {
        var authToken = localStorage.getItem("Authorization") || "";
        const responseTotal = await fetch(Endpoints.Testbase + Endpoints.Responses + Endpoints.ResponsesActiveOnly, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (responseTotal.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const formsTotal = await responseTotal.json();
        setTotal(formsTotal.length);
    }

    async function getTodayTrends() {
        var authToken = localStorage.getItem("Authorization") || "";
        const responseToday= await fetch(Endpoints.Testbase + Endpoints.Responses + Endpoints.ResponsesTodayOnly, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (responseToday.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const formsToday = await responseToday.json();
        // map hour and form response counts
        var trendData = new Map<number, number>();
        formsToday.forEach(function(form: any) {
            var date = new Date(form.createdAt);
            var hour = date.getHours();
            if (trendData.has(hour)) {
                trendData.set(hour, (trendData.get(hour) || 0) + 1);
            } else {
                trendData.set(hour, 1);
            }
        });
        // convert map to array for recharts line chart
        var trendDataArray: Array<ChartData> = [];
        for (var i = 0; i < 24; i++) {
            trendDataArray.push({index: i, responses: trendData.get(i) || 0});
        }
        setTodayTrends(trendDataArray);
        setToday(formsToday.length);
    }

    useEffect(() => {
        getTotal();
        getTodayTrends();
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
                <DashboardChart data={todayTrends}></DashboardChart>
                <div className="chart-stats-cards">
                    <StatCard title={Constants.Total} stat={total}></StatCard>
                    <StatCard title={Constants.Today} stat={today}></StatCard>
                </div>
            </div>
        </div>
    );
}