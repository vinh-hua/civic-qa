import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, InquiryData } from '../Components/SubDashboard';
import { FormInquiryView } from './FormInquiryView'; 
import * as Endpoints from '../Constants/Endpoints';

export type ResponsesProps = {
    header?: string;
    subjectTitle?: string;
    data: Array<InquiryData>;
    hideInquiryBackArrow?: boolean;
}

export function Inquiries(props: ResponsesProps) {
    const headerTitle = props.header || "Form Inquiries";
    const subjecTitle = props.subjectTitle || "CURRENT INQUIRIES";
    const hideBackArrow = props.hideInquiryBackArrow || false;
    const [onInquiryView, setInquiryView] = useState(false);
    const [responseData, setData] = useState<InquiryData[]>([]);
    const [specificInquiryData, setSpecificInquiryData] = useState<InquiryData>();

    const getInquiries = async() => {
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
        var formInquiries: Array<InquiryData> = [];
        forms.forEach(function(inquiry: any) {
            var d = new Date(inquiry.createdAt);
            var t = d.toLocaleString("en-US");
            formInquiries.push({id: inquiry.id, email: inquiry.emailAddress, name: inquiry.name + " / " + inquiry.subject, value: t, body: inquiry.body});
        });
        setData(formInquiries);
    }

    useEffect(() => {
        getInquiries();
    }, []);

    function setSpecificInquiryContent(formResponse: InquiryData) {
        setInquiryView(true);
        setSpecificInquiryData(formResponse);
    }

    function setSpecificView() {
        setInquiryView(false);
        getInquiries();
    }
    
    return (
        <div className="dashboard sub-dashboard">
            {onInquiryView ? 
            <FormInquiryView responseId={specificInquiryData?.id || ""} email={specificInquiryData?.email || ""} title="Inquiry" subject={specificInquiryData?.name || ""} body={specificInquiryData?.body || ""} setSpecificView={setSpecificView} hideBackArrow={hideBackArrow}></FormInquiryView> :
            <div>
                <Header title={headerTitle}></Header>
                <SubDashboard title={subjecTitle} data={props.data ? props.data : responseData} changeViewFunc={setSpecificInquiryContent} fullPageView={true} subHeaderValue={props.data ? props.data.length : responseData.length}></SubDashboard>
            </div>}
        </div>
    );
}