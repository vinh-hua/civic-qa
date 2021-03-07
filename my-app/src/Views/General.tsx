import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Responses } from './Responses';

// TODO: will data be pre-sorted on back-end?
// currently using test data
function getInitialSubDashboardData(): Array<SubDashboardData> {
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

function getSpecificData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    return data as Array<SubDashboardData>;
}

export function General() {
    const testData = getInitialSubDashboardData();
    const testSpecificData = getSpecificData();
    const [onSpecificView, setSpecificView] = useState(false);

    let statCards = [
        {title: "New Today", stat: 288},
        {title: "This Week", stat: 106},
        {title: "Topics", stat: 24}
    ];

    function specificView() {
        setSpecificView(true);
    }

    function initialView() {
        setSpecificView(false);
    }

    return (
        onSpecificView ? 
        <div> 
            <div className="dashboard sub-dashboard">
                <button className="exit-button" onClick={initialView}><img src="./assets/icons/back-arrow.png"></img></button>
            </div>
            <Responses header="General Inquiries" data={testSpecificData}></Responses>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="General Inquiries"></Header>
                <SubDashboard title="TOP SUBJECTS" data={testData} changeViewFunc={specificView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                    <StatCardRow cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}