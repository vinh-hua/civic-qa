import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Responses } from './Responses';

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

function getSpecificData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    data.push({name: "Test", value: 123});
    return data as Array<SubDashboardData>;
}

export function Casework() {
    const testData = getSubDashboardData();
    const testSpecificData = getSpecificData();
    const [onSpecificView, setSpecificView] = useState(false);

    let statCards = [
        {title: "New Today", stat: 9},
        {title: "This Week", stat: 23},
        {title: "Topics", stat: 10}
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
            <Responses header="Casework" data={testSpecificData}></Responses>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="Casework"></Header>
                <SubDashboard title="TOPIC" data={testData} changeViewFunc={specificView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                    <StatCardRow cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}