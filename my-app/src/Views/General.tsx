import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { StatCardRow } from '../Components/StatCardRow';
import { Inquiries } from './Inquiries';
import { useSelector } from 'react-redux';
import { AppState } from '../Redux/Reducers/rootReducer'
import * as Endpoints from '../Constants/Endpoints';

export function General() {
    const { auth } = useSelector((state: AppState) => state.auth);
    const [onInquiriesView, setInquiriesView] = useState(false);
    const [specificViewTitle, setSpecificViewTitle] = useState("");
    const [specificSubjectData, setSpecificSubjectData] = useState<SubDashboardData[]>([]);
    const [subjectsResponsesData, setSubjectsResponsesData] = useState<Map<string, SubDashboardData[]>>();
    const [subjectsInquiries, setSubjectsInquiries] = useState<SubDashboardData[]>([]);
    const [summaryToday, setSummaryToday] = useState(0);
    const [summaryTotal, setSummaryTotal] = useState(0);
    const [summaryTopics, setSummaryTopics] = useState(0);

    const getResponses = async() => {
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
        const responsesGeneral = await response.json();
        var formResponses: SubDashboardData[] = [];
        let subjectsMap = new Map<string, SubDashboardData[]>();
        let subjectsInquiries = new Map<string, number>();

        responsesGeneral.forEach(function(formResponse: any) {
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            var subjects = formResponse.tags;
            var data: SubDashboardData = {id: formResponse.id, email: formResponse.emailAddress, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body}

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
            var subText = " inquiry";
            if ((subjectsInquiries.get(key) || 0) > 1) {
                subText = " inquiries";
            }
            inquiries.push({name: key, value: subjectsInquiries.get(key) + subText});
        });

        inquiries.sort((a, b) => (a.value > b.value) ? -1 : (a.value === b.value) ? -1 : 1);

        setSummaryTopics(inquiries.length);
        setSummaryTotal(formResponses.length);
        setSubjectsInquiries(inquiries);
        setSubjectsResponsesData(subjectsMap);
    }

    const getResponsesToday = async() => {
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
        const responsesToday = await response.json();
        setSummaryToday(responsesToday.length);
    }

    function inquiriesView(data: SubDashboardData) {
        setSpecificViewTitle(data.name);
        setSpecificSubjectData(subjectsResponsesData?.get(data.name) || []);
        setInquiriesView(true);
    }

    function initialView() {
        setInquiriesView(false);
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
        onInquiriesView ? 
        <div> 
            <div className="dashboard sub-dashboard">
                <button className="exit-button" onClick={initialView}><img className="back-arrow" src="./assets/icons/arrow.svg"></img></button>
            </div>
            <Inquiries header="General" subjectTitle={specificViewTitle} data={specificSubjectData} hideInquiryBackArrow={true}></Inquiries>
        </div>
        : <div className="dashboard sub-dashboard">
            <div>
                <Header title="General Inquiries"></Header>
                <SubDashboard title="TOPICS" data={subjectsInquiries} changeViewFunc={inquiriesView} emailTemplates={false} fullPageView={false}></SubDashboard>
                <div className="sub-summary">
                    <SubHeaderLine title="SUMMARY" subHeaderValue={"Active Inquiries"}></SubHeaderLine>
                    <StatCardRow spaceEven={false} cards={statCards}></StatCardRow>
                </div>
            </div>
        </div>
    );
}