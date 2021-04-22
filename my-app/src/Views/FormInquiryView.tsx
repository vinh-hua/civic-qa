import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { Tags } from '../Components/Tags';
import * as Endpoints from '../Constants/Endpoints';
import "./FormInquiryView.css";

export type FormResponseViewProps = {
    responseId: string;
    email: string;
    title: string;
    subject: string;
    body: string;
    setSpecificView: Function;
    hideBackArrow: boolean;
};

export function FormInquiryView(props: FormResponseViewProps) {
    const [isResolved, setIsResolved] = useState(true);
    const [tags, setTags] = useState<any[]>([]);
    const [messageResponse, setMessageResponse] = useState("");

    async function createMailto() {
        var mailtoRequest = {to: [props.email], subject: props.subject, body: messageResponse};
        var jsonMailtoRequest = JSON.stringify(mailtoRequest);
        const response = await fetch(Endpoints.Base + Endpoints.Mailto, {
            method: "POST",
            body: jsonMailtoRequest,
            headers: new Headers({
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            console.log("Error creating mailto");
            return;
        }
        const mailtoString = await response.text();
        window.location.href = mailtoString;
    }

    const getTags = async() => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Base + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retreiving response tags");
            return;
        }
        const tags = await response.json();
        var tagList: any[] = [];
        tags.forEach((tag: any) => {
            tagList.push(tag.value);
        });
        setTags(tagList);
    }

    async function removeTag(tagValue: string) {
        var authToken = localStorage.getItem("Authorization") || "";
        var tagJson = JSON.stringify({value: tagValue});
        const response = await fetch(Endpoints.Base + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
            method: "DELETE",
            body: tagJson,
            headers: new Headers({
                "Authorization": authToken,
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            console.log("Error creating tag");
            return;
        }
        getTags();
    }

    async function addTag(tagValue: string) {
        var authToken = localStorage.getItem("Authorization") || "";
        var tagJson = JSON.stringify({value: tagValue});
        const response = await fetch(Endpoints.Base + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
            method: "POST",
            body: tagJson,
            headers: new Headers({
                "Authorization": authToken,
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            console.log("Error creating tag");
            return;
        }
        getTags();
    }

    const resolveResponse = async(id: string, isResolved: boolean) => {
        var authToken = localStorage.getItem("Authorization") || "";
        var patchActive = JSON.stringify({active: !isResolved});
        const response = await fetch(Endpoints.Base + Endpoints.Responses + "/" + id, {
            method: "PATCH",
            body: patchActive,
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error marking form response as resolved");
            return;
        }
    }

    async function clickCheckbox() {
        setIsResolved(isResolved => !isResolved);
        resolveResponse(props.responseId, isResolved);
    }

    useEffect(() => {
        getTags();
    }, []);

    return(
        <div>
            {props.hideBackArrow ? null : <button className="exit-button" onClick={() => props.setSpecificView()}><img className="back-arrow" src="./assets/icons/arrow.svg"></img></button>}
            <Header title={props.title}></Header>
            <SubHeaderLine title={props.subject}></SubHeaderLine>
            <Tags addTag={addTag} removeTag={removeTag} values={tags}></Tags>
            <div className="form-response-container">
                <div className="form-response">
                    <p className="form-response-body">{props.body}</p>
                    <textarea className="form-response-message" value={messageResponse} onChange={e => setMessageResponse(e.target.value)}></textarea>
                    <div className="resolved-send-container">
                        <label className="resolved-label" >
                            <input id="resolved-check-box" className="resolved-check-box" type="checkbox" onClick={() => clickCheckbox()}></input>
                            Resolved
                        </label>
                        <button className="send-btn" placeholder="Message" onClick={createMailto}>Send</button>
                    </div>
                </div>
            </div>
        </div>
    );
}