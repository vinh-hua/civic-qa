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
    data.push({name: "Test", value: 22});
    return data as Array<SubDashboardData>;
}

export function Casework() {
    const testData = getSubDashboardData();
    const [onSpecificView, setSpecificView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificTopicData, setSpecificTopicData] = useState<SubDashboardData[]>([]);
    const [summaryToday, setSummaryToday] = useState(0);
    const [summaryWeek, setSummaryWeek] = useState(0);
    const [summaryTopics, setSummaryTopics] = useState(0);


    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "?" + Endpoints.ResponsesActiveCasework, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const responses = await response.json();
        var formResponses: Array<SubDashboardData> = [];
        responses.forEach(function(formResponse: any) {
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            formResponses.push({id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body});
        });
        setSpecificTopicData(formResponses);
    }

    const getResponsesToday = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "?" + Endpoints.ResponsesActiveCasework + "&" + Endpoints.ResponsesTodayOnly, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const responsesToday = await response.json();
        setSummaryToday(responsesToday.length);
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
        getResponsesToday();
    }, []);

    let statCards = [
        {title: "New Today", stat: summaryToday},
        {title: "This Week", stat: summaryWeek},
        {title: "Topics", stat: summaryTopics}
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