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
    const [data, setData] = useState(test_data);
    const [onSpecificView, setSpecificView] = useState(false);

    let statCards = [
        {title: "New Today", stat: 9},
        {title: "Pending", stat: 23},
        {title: "Unanswered", stat: 55}
    ];

    return (
        <div className="dashboard sub-dashboard">
            <Header title="Casework"></Header>
            <SubDashboard title="TOPIC" data={data} setData={setData} emailTemplates={false} hasRespondOption={false}></SubDashboard>
            <div className="sub-summary">
                <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                <StatCardRow cards={statCards}></StatCardRow>
            </div>
        </div>
    );
}