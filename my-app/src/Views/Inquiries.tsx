import React, { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { FormInquiryView } from './FormInquiryView'; 
import * as Endpoints from '../Constants/Endpoints';

export type ResponsesProps = {
    header?: string;
    subjectTitle?: string;
    data: Array<SubDashboardData>;
    hideInquiryBackArrow?: boolean;
}

export function Inquiries(props: ResponsesProps) {
    const headerTitle = props.header || "Form Inquiries";
    const subjecTitle = props.subjectTitle || "CURRENT INQUIRIES";
    const hideBackArrow = props.hideInquiryBackArrow || false;
    const [onResponseView, setResponseView] = useState(false);
    const [responseData, setData] = useState<SubDashboardData[]>([]);
    const [specificResponseData, setSpecificResponseData] = useState<SubDashboardData>();

    const getResponses = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Base + Endpoints.ResponsesActiveOnly, {
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
            formResponses.push({id: formResponse.id, email: formResponse.emailAddress, name: formResponse.name + " / " + formResponse.subject, value: t, body: formResponse.body});
        });
        setData(formResponses);
    }

    useEffect(() => {
        getResponses();
    }, []);

    function setSpecificResponseContent(formResponse: SubDashboardData) {
        setResponseView(true);
        setSpecificResponseData(formResponse);
    }

    function setSpecificView() {
        setResponseView(false);
        getResponses();
    }
    
    return (
        <div className="dashboard sub-dashboard">
            {onResponseView ? 
            <FormInquiryView responseId={specificResponseData?.id || ""} email={specificResponseData?.email || ""} title="Inquiry" subject={specificResponseData?.name || ""} body={specificResponseData?.body || ""} setSpecificView={setSpecificView} hideBackArrow={hideBackArrow}></FormInquiryView> :
            <div>
                <Header title={headerTitle}></Header>
                <SubDashboard title={subjecTitle} data={props.data ? props.data : responseData} changeViewFunc={setSpecificResponseContent} emailTemplates={false} fullPageView={true} subHeaderValue={props.data ? props.data.length : responseData.length}></SubDashboard>
            </div>}
        </div>
    );
}