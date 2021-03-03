import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';

// TODO: will data be pre-sorted on back-end?
// currently using test data
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "Employment Card", value: 22});
    data.push({name: "Citizenship", value: 21});
    data.push({name: "Social Security #", value: 18});
    data.push({name: "Tax Documents", value: 15});
    data.push({name: "Other", value: 11});
    return data as Array<SubDashboardData>;
}

export function Casework() {
    const test_data = getSubDashboardData();
    const [onSpecificView, setSpecificView] = useState(false);

    let statCards = [
        {title: "New Today", stat: 9},
        {title: "This Week", stat: 23},
        {title: "Topics", stat: 10}
    ];

    return (
        <div className="dashboard sub-dashboard">
            <Header title="Casework"></Header>
            <SubDashboard title="TOPIC" data={test_data} setSpecificView={setSpecificView} emailTemplates={false} fullPageView={false} hasRespondOption={false} viewButton={false}></SubDashboard>
            <div className="sub-summary">
                <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                <StatCardRow cards={statCards}></StatCardRow>
            </div>
        </div>
    );
}