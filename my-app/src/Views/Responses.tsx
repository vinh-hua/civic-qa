import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';

// currently using test data
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "Lucille Harmon / COVID-19 Stimulus", value: "4:35pm"});
    data.push({name: "Vincent Donahue / HB 1009", value: "3:56pm"});
    data.push({name: "Taylor Todd / COVID-19 Vaccination", value: "3:02pm"});
    data.push({name: "Serena Gomez / SPD Reform", value: "2:47pm"});
    data.push({name: "Anna Nightingale / SB 3420", value: "2:47pm"});
    data.push({name: "Lara Cooke / Employment Card", value: "2:47pm"});
    data.push({name: "Tax Documents", value: "2:47pm"});
    data.push({name: "Dante Preston / Homelessness", value: "2:47pm"});
    data.push({name: "Jordan Phan / COVID-19 Vaccination", value: "2:47pm"});
    data.push({name: "Brooklyn Drake / Citizenship", value: "2:47pm"});
    return data as Array<SubDashboardData>;
}

export function Responses() {
    const test_data = getSubDashboardData();
    const [data, setData] = useState(test_data);
    return (
        <div className="dashboard sub-dashboard">
            <Header title="Form Responses"></Header>
            <SubDashboard title="CURRENT RESPONSES" data={data} setData={setData} emailTemplates={false} formResponses={true} hasRespondOption={false} viewButton={false} subHeaderNumber={342}></SubDashboard>
        </div>
    );
}