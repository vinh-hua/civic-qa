import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, InquiryData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Inquiries } from './Inquiries';
import * as Endpoints from '../Constants/Endpoints';

export function General() {
    const auth = localStorage.getItem("Authorization") || "";
    const [onInquiriesView, setInquiriesView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificTopicData, setSpecificTopicData] = useState<InquiryData[]>([]);
    const [topicsData, setTopicsData] = useState<Map<string, InquiryData[]>>();
    const [topicsInquiries, setTopicsInquiries] = useState<InquiryData[]>([]);
    const [summaryToday, setSummaryToday] = useState(0);
    const [summaryTotal, setSummaryTotal] = useState(0);
    const [summaryTopics, setSummaryTopics] = useState(0);

    const getInquiries = async() => {
        const response = await fetch(Endpoints.Base + Endpoints.ResponsesActiveGeneral, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const inquiriesGeneral = await response.json();
        var formInquiries: InquiryData[] = [];
        let subjectsMap = new Map<string, InquiryData[]>();
        let subjectsInquiries = new Map<string, number>();

        inquiriesGeneral.forEach(function(inquiry: any) {
            var d = new Date(inquiry.createdAt);
            var t = d.toLocaleString("en-US");
            var subjects = inquiry.tags;
            var data: InquiryData = {id: inquiry.id, email: inquiry.emailAddress, name: inquiry.name + " / " + inquiry.subject, value: t, body: inquiry.body}

            subjects.forEach((subject: any) => {
                if (subjectsMap.has(subject.value)) {
                    var getList = subjectsMap.get(subject.value);
                    getList?.push(data);
                    subjectsMap.set(subject.value, getList || []);
                } else {
                    var newList: InquiryData[] = [];
                    newList.push(data);
                    subjectsMap.set(subject.value, newList);
                }

                subjectsInquiries.set(subject.value, (subjectsInquiries.get(subject.value) || 0) + 1);

            });
            formInquiries.push(data);
        });

        var inquiries: InquiryData[] = [];
        Array.from(subjectsInquiries.keys()).forEach((key) => {
            var subText = " inquiry";
            if ((subjectsInquiries.get(key) || 0) > 1) {
                subText = " inquiries";
            }
            inquiries.push({name: key, value: subjectsInquiries.get(key) + subText});
        });

        inquiries.sort((a, b) => (a.value > b.value) ? -1 : (a.value === b.value) ? -1 : 1);

        setSummaryTopics(inquiries.length);
        setSummaryTotal(formInquiries.length);
        setTopicsInquiries(inquiries);
        setTopicsData(subjectsMap);
    }

    const getInquiriesToday = async() => {
        const response = await fetch(Endpoints.Base +  Endpoints.ResponsesActiveGeneralTodayOnly, {
            method: "GET",
            headers: new Headers({
                "Authorization": auth
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const inquiriesToday = await response.json();
        setSummaryToday(inquiriesToday.length);
    }

    function inquiriesView(data: InquiryData) {
        setSpecificViewTitle(data.name);
        setSpecificTopicData(topicsData?.get(data.name) || []);
        setInquiriesView(true);
    }

    function initialView() {
        setInquiriesView(false);
        getInquiries();
    }
    
    useEffect(() => {
        getInquiries();
        getInquiriesToday();
    }, []);


    let statCards = [
        {title: "New Today", stat: summaryToday},
        {title: "Total", stat: summaryTotal},
        {title: "Topics", stat: summaryTopics}
    ];

    return (
        onInquiriesView ? 
        <div> 
            <div className="dashboard sub-dashboard">
                <button className="exit-button" onClick={initialView}><img className="back-arrow" src="./assets/icons/arrow.svg"></img></button>
            </div>
            <Inquiries header="General" subjectTitle={specificViewTitle} data={specificTopicData} hideInquiryBackArrow={true}></Inquiries>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="General Topics"></Header>
                <SubDashboard title="TOPICS" data={topicsInquiries} changeViewFunc={inquiriesView} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY" subHeaderValue={"Active Inquiries"}></SubHeaderLine>
                    <StatCardRow spaceEven={false} cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}