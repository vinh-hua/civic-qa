import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';

// TODO: will data be pre-sorted on back-end?
// currently using test data
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "SPD - Reform", value: 123});
    data.push({name: "COVID-19 - Stimulus", value: 119});
    data.push({name: "Homelessness- Shelter", value: 77});
    data.push({name: "Investments in BIPOC Communities", value: 62});
    data.push({name: "SPD - Accountability", value: 36});
    data.push({name: "SPD - Accountability", value: 36});
    data.push({name: "SPD - Accountability", value: 36});
    data.push({name: "SPD - Accountability", value: 36});
    data.push({name: "Other", value: 52});
    return data as Array<SubDashboardData>;
}

export function General() {
    const test_data = getSubDashboardData();
    const [data, setData] = useState(test_data);
    const [onSpecificView, setSpecificView] = useState(false);

    let statCards = [
        {title: "New Today", stat: 288},
        {title: "Follow-Ups", stat: 106},
        {title: "Unanswered", stat: 934}
    ];

    return (
        <div className="dashboard sub-dashboard">
            <Header title="General Inquiries"></Header>
            <SubDashboard title="Top Subjects" data={data} setData={setData} hasRespondOption={true}></SubDashboard>
            <div className="sub-summary">
                <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                <StatCardRow cards={statCards}></StatCardRow>
            </div>
        </div>
    );
}