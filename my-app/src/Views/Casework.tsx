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
    const [specificTopicsData, setSpecificTopicsData] = useState<SubDashboardData[]>([]);
    const [topicsResponsesData, setTopicsResponsesData] = useState<Map<string, SubDashboardData[]>>();
    const [topicsCases, setTopicsCases] = useState<SubDashboardData[]>([]);
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
        const responsesGeneral = await response.json();
        var formResponses: SubDashboardData[] = [];
        let topicsMap = new Map<string, SubDashboardData[]>();
        let topicsCases = new Map<string, number>();

        responsesGeneral.forEach(function(formResponse: any) {
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            var topics = formResponse.tags;
            var data: SubDashboardData = {id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body}

            topics.forEach((topic: any) => {
                if (topicsMap.has(topic.value)) {
                    var getList = topicsMap.get(topic.value);
                    getList?.push(data);
                    topicsMap.set(topic.value, getList || []);
                } else {
                    var newList: SubDashboardData[] = [];
                    newList.push(data);
                    topicsMap.set(topic.value, newList);
                }

                topicsCases.set(topic.value, (topicsCases.get(topic.value) || 0) + 1);

            });
            formResponses.push(data);
        });

        var cases: SubDashboardData[] = [];
        Array.from(topicsCases.keys()).forEach((key) => {
            cases.push({name: key, value: topicsCases.get(key)});
        })

        setTopicsCases(cases);
        setTopicsResponsesData(topicsMap);
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
        setSpecificTopicsData(topicsResponsesData?.get(data.name) || []);
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
            <Responses header="Casework" subjectTitle={specificViewTitle} data={specificTopicsData}></Responses>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="Casework"></Header>
                <SubDashboard title="TOPIC" data={topicsCases} changeViewFunc={specificView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                    <StatCardRow spaceEven={false} cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}