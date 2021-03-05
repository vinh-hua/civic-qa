import React, { useState } from 'react';
import { Header } from '../Components/Header';
import { SubDashboard, SubDashboardData } from '../Components/SubDashboard';
import { SpecificView } from '../Components/SpecificView'; 

// currently using test data
// TODO: sort by time ?? backend feature??
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "Lucille Harmon / COVID-19 Stimulus", value: "4:35pm"});
    data.push({name: "Vincent Donahue / HB 1009", value: "3:56pm"});
    data.push({name: "Taylor Todd / COVID-19 Vaccination", value: "3:02pm"});
    data.push({name: "Serena Gomez / SPD Reform", value: "2:47pm"});
    data.push({name: "Anna Nightingale / SB 3420", value: "2:47pm"});
    data.push({name: "Lara Cooke / Employment Card", value: "2:47pm"});
    data.push({name: "Tax Documents", value: "2:47pm"});
    data.push({name: "Dante Preston / Homelessness", value: "2:47pm"});
    data.push({name: "Jordan Phan / COVID-19 Vaccination", value: "2:47pm"});
    data.push({name: "Brooklyn Drake / Citizenship", value: "2:47pm"});
    return data as Array<SubDashboardData>;
}

export function Responses() {
    const test_data = getSubDashboardData();
    const test_body = "Dear WA 36th Legislative Staff, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam maximus diam egestas augue dignissim, quis accumsan tortor pulvinar. Suspendisse mattis quam magna, ut dapibus leo volutpat non. Donec sapien mauris, semper non odio at, gravida posuere massa. Sed mattis diam id sapien semper sodales. Nam in justo ultrices, facilisis arcu vitae, ornare velit. Nam vitae aliquam... More text to test overflow, is it overflowing, is it overflowing, is it overflowing, is it overflowing, is it overflowing, is it overflowing, is it overflowing?";
    const [onResponseView, setResponseView] = useState(false);
    const [responseSubject, setResponseSubject] = useState("");
    const [responseBody, setResponseBody] = useState("");

    function setResponseContent(subject: string) {
        setResponseView(true);
        setResponseSubject(subject);
        setResponseBody(test_body);
    }
    
    return (
        <div className="dashboard sub-dashboard">
            {onResponseView? <SpecificView title="Form Responses" subject={responseSubject} body={responseBody} subHeaderNumber={342} setSpecificView={() => setResponseView(false)}></SpecificView> :
            <div>
                <Header title="Form Responses"></Header>
                <SubDashboard title="CURRENT RESPONSES" data={test_data} setSpecificView={setResponseContent} emailTemplates={false} fullPageView={true} hasRespondOption={false} viewButton={false} subHeaderNumber={342}></SubDashboard>
            </div>}
        </div>
    );
}