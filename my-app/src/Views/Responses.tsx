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
    const [responseData, setData] = useState<SubDashboardData[]>([]);
    const [specificResponseData, setSpecificResponseData] = useState<SubDashboardData>();

    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch("http://localhost/v0/responses?activeOnly=true", {
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
            var d = new Date(formResponse.createdAt);
            var t = d.toLocaleString("en-US");
            formResponses.push({id: formResponse.id, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body});
        });
        setData(formResponses);
    }

    useEffect(() => {
        getResponses();
    }, []);

    function setResponseContent(formResponse: SubDashboardData) {
        setResponseView(true);
        setSpecificResponseData(formResponse);
    }
    
    return (
        <div className="dashboard sub-dashboard">
            {onResponseView ? <FormResponseView responseId={specificResponseData?.id || ""} title="Form Responses" subject={specificResponseData?.name || ""} body={specificResponseData?.body || ""} setSpecificView={() => setResponseView(false)}></FormResponseView> :
            <div>
                <Header title={headerTitle}></Header>
                <SubDashboard title={subjecTitle} data={responseData} changeViewFunc={setResponseContent} emailTemplates={false} fullPageView={true} subHeaderNumber={responseData.length}></SubDashboard>
            </div>}
        </div>
    );
}