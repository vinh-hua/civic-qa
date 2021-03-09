import React, { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Responses } from './Responses';
import * as Endpoints from '../Constants/Endpoints';

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
    const testData = getSubDashboardData();
    const [onSpecificView, setSpecificView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificTopicData, setSpecificTopicData] = useState<SubDashboardData[]>([]);

    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + Endpoints.ResponsesActiveCasework, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const responsesGeneral = await response.json();
        var formResponses: Array<SubDashboardData> = [];
        responsesGeneral.forEach(function(formResponse: any) {
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            formResponses.push({id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body});
        });
        setSpecificTopicData(formResponses);
    }

    function specificView(data: SubDashboardData) {
        setSpecificViewTitle(data.name);
        setSpecificView(true);
    }

    function initialView() {
        setSpecificView(false);
    }
    
    useEffect(() => {
        getResponses();
    }, []);

    let statCards = [
        {title: "New Today", stat: 9},
        {title: "This Week", stat: 23},
        {title: "Topics", stat: 10}
    ];

    return (
        onSpecificView ? 
        <div> 
            <div className="dashboard sub-dashboard">
                <button className="exit-button" onClick={initialView}><img src="./assets/icons/back-arrow.png"></img></button>
            </div>
            <Responses header="Casework" subjectTitle={specificViewTitle} data={specificTopicData}></Responses>
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