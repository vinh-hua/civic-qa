import React, { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Responses } from './Responses';
import { useSelector } from 'react-redux';
import { AppState } from '../Redux/Reducers/rootReducer'
import * as Endpoints from '../Constants/Endpoints';

export function Casework() {
    const { auth } = useSelector((state: AppState) => state.auth);
    const [onSpecificView, setSpecificView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificTopicsData, setSpecificTopicsData] = useState<SubDashboardData[]>([]);
    const [topicsResponsesData, setTopicsResponsesData] = useState<Map<string, SubDashboardData[]>>();
    const [topicsCases, setTopicsCases] = useState<SubDashboardData[]>([]);
    const [summaryToday, setSummaryToday] = useState(0);
    const [summaryTotal, setSummaryTotal] = useState(0);
    const [summaryTopics, setSummaryTopics] = useState(0);


    const getResponses = async() => {
        const response = await fetch(Endpoints.Base + Endpoints.ResponsesActiveCasework, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
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
            var data: SubDashboardData = {id: formResponse.id, email: formResponse.emailAddress, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body}

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
            var subText = " case";
            if ((topicsCases.get(key) || 0) > 1) {
                subText = " cases";
            }
            cases.push({name: key, value: topicsCases.get(key) + subText});
        })

        cases.sort((a, b) => (a.value > b.value) ? -1 : (a.value === b.value) ? -1 : 1);

        setSummaryTopics(cases.length);
        setSummaryTotal(formResponses.length);
        setTopicsCases(cases);
        setTopicsResponsesData(topicsMap);
    }

    const getResponsesToday = async() => {
        const response = await fetch(Endpoints.Base + Endpoints.Responses + "?" + Endpoints.ResponsesActiveCasework + "&" + Endpoints.ResponsesTodayOnly, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
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
        {title: "Total", stat: summaryTotal},
        {title: "Topics", stat: summaryTopics}
    ];

    return (
        onSpecificView ? 
        <div> 
            <div className="dashboard sub-dashboard">
                <button className="exit-button" onClick={initialView}><img className="back-arrow" src="./assets/icons/arrow.svg"></img></button>
            </div>
            <Responses header="Casework" subjectTitle={specificViewTitle} data={specificTopicsData}></Responses>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="Casework"></Header>
                <SubDashboard title="TOPIC" data={topicsCases} changeViewFunc={specificView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY" subHeaderValue={"Active Cases"}></SubHeaderLine>
                    <StatCardRow spaceEven={false} cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}