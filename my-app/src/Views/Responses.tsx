import React, { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { FormResponseView } from './FormResponseView'; 

export type ResponsesProps = {
    header?: string;
    subjectTitle?: string;
    data: Array<SubDashboardData>;
}

export function Responses(props: ResponsesProps) {
    const headerTitle = props.header || "Form Responses";
    const subjecTitle = props.subjectTitle || "CURRENT RESPONSES";
    const [onResponseView, setResponseView] = useState(false);
    const [responseSubject, setResponseSubject] = useState("");
    const [responseBody, setResponseBody] = useState("");
    const [data, setData] = useState<SubDashboardData[]>([]);

    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch("http://localhost/v0/responses", {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retrieving form responses");
            return;
        }
        const forms = await response.json();
        var formResponses: Array<SubDashboardData> = [];
        forms.forEach(function(formResponse: any) {
            formResponses.push({id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: formResponse.createdAt, body: formResponse.body});
        });
        setData(formResponses);
    }

    useEffect(() => {
        getResponses();
    }, []);

    function setResponseContent(formResponse: SubDashboardData) {
        setResponseView(true);
        setResponseSubject(formResponse.name);
        setResponseBody(formResponse.body || "");
    }
    
    return (
        <div className="dashboard sub-dashboard">
            {onResponseView ? <FormResponseView responseId={"1"} title="Form Responses" subject={responseSubject} body={responseBody} setSpecificView={() => setResponseView(false)}></FormResponseView> :
            <div>
                <Header title={headerTitle}></Header>
                <SubDashboard title={subjecTitle} data={data} changeViewFunc={setResponseContent} emailTemplates={false} fullPageView={true} subHeaderNumber={data.length}></SubDashboard>
            </div>}
        </div>
    );
}