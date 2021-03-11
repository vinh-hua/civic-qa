import React, { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Responses } from './Responses';
import * as Endpoints from '../Constants/Endpoints';

export function General() {
    const [onSpecificView, setSpecificView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificSubjectData, setSpecificSubjectData] = useState<SubDashboardData[]>([]);
    const [subjectsResponsesData, setSubjectsResponsesData] = useState<Map<string, SubDashboardData[]>>();
    const [subjectsInquiries, setSubjectsInquiries] = useState<SubDashboardData[]>([]);
    const [summaryToday, setSummaryToday] = useState(0);
    const [summaryTotal, setSummaryTotal] = useState(0);
    const [summaryTopics, setSummaryTopics] = useState(0);

    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "?" + Endpoints.ResponsesActiveGeneral, {
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
        let subjectsMap = new Map<string, SubDashboardData[]>();
        let subjectsInquiries = new Map<string, number>();

        responsesGeneral.forEach(function(formResponse: any) {
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            var subjects = formResponse.tags;
            var data: SubDashboardData = {id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body}

            subjects.forEach((subject: any) => {
                if (subjectsMap.has(subject.value)) {
                    var getList = subjectsMap.get(subject.value);
                    getList?.push(data);
                    subjectsMap.set(subject.value, getList || []);
                } else {
                    var newList: SubDashboardData[] = [];
                    newList.push(data);
                    subjectsMap.set(subject.value, newList);
                }

                subjectsInquiries.set(subject.value, (subjectsInquiries.get(subject.value) || 0) + 1);

            });
            formResponses.push(data);
        });

        var inquiries: SubDashboardData[] = [];
        Array.from(subjectsInquiries.keys()).forEach((key) => {
            inquiries.push({name: key, value: subjectsInquiries.get(key) + " inquiries"});
        });

        inquiries.sort((a, b) => (a.value > b.value) ? -1 : (a.value === b.value) ? -1 : 1);

        setSummaryTopics(inquiries.length);
        setSummaryTotal(formResponses.length);
        setSubjectsInquiries(inquiries);
        setSubjectsResponsesData(subjectsMap);
    }

    const getResponsesToday = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "?" + Endpoints.ResponsesActiveGeneral + "&" + Endpoints.ResponsesTodayOnly, {
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
        setSpecificSubjectData(subjectsResponsesData?.get(data.name) || []);
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
                <button className="exit-button" onClick={initialView}><img src="./assets/icons/back-arrow.png"></img></button>
            </div>
            <Responses header="General Inquiries" subjectTitle={specificViewTitle} data={specificSubjectData}></Responses>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="General Inquiries"></Header>
                <SubDashboard title="TOP SUBJECTS" data={subjectsInquiries} changeViewFunc={specificView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY"></SubHeaderLine>
                    <StatCardRow spaceEven={false} cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}