import { useEffect, useState } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { Tag } from '../Components/Tag';
import { Tags } from '../Components/Tags';
import * as Endpoints from '../Constants/Endpoints';
import "./FormResponseView.css";

export type FormResponseViewProps = {
    responseId: string;
    title: string;
    subject: string;
    body: string;
    setSpecificView: Function;
};

export function FormResponseView(props: FormResponseViewProps) {
    const [isResolved, setIsResolved] = useState(true);
    const [tags, setTags] = useState<any[]>([]);
    const [messageResponse, setMessageResponse] = useState("");

    async function createMailto() {
        var mailtoRequest = {to: ["test@test.com"], subject: props.subject, body: messageResponse};
        var jsonMailtoRequest = JSON.stringify(mailtoRequest);
        const response = await fetch(Endpoints.Testbase + Endpoints.Mailto, {
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
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
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
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
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
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + props.responseId + Endpoints.ResponsesTags, {
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
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + id, {
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
            <button className="exit-button" onClick={() => props.setSpecificView()}><img src="./assets/icons/back-arrow.png"></img></button>
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