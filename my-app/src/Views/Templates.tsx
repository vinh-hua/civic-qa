import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import './Templates.css';

function getTemplatesData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "COVID-19 Vaccination", value: "Dear ____, Current COVID-19..."});
    data.push({name: "Voter Registration Info", value: "Dear ____, Instructions on how..."});
    data.push({name: "Voter Registration Info", value: "Dear ____, Instructions on how..."});
    data.push({name: "Voter Registration Info", value: "Dear ____, Instructions on how..."});
    data.push({name: "Voter Registration Info", value: "Dear ____, Instructions on how..."});
    data.push({name: "Voter Registration Info", value: "Dear ____, Instructions on how..."});
    return data as Array<SubDashboardData>;
}

export function Templates() {
    const test_data = getTemplatesData();
    const [data, setData] = useState(test_data);
    const [onSpecificView, setSpecificView] = useState(false);

    return (
        <div className="dashboard sub-dashboard">
            <Header title="Email Templates"></Header>
            <SubDashboard title="Templates" data={data} setData={setData} emailTemplates={true} formResponses={false} hasRespondOption={false} viewButton={true}></SubDashboard>
            <button className="templates-btn">Create New</button>
        </div>
    );
}