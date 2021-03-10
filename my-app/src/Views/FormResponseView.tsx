import { useState } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from '../Components/SubHeaderLine';
import { Tag } from '../Components/Tag';
import { TagAdd } from '../Components/TagAdd';
import * as Endpoints from '../Constants/Endpoints';
import "./FormResponseView.css";

export type FormResponseViewProps = {
    responseId: string;
    title: string;
    subject: string;
    body: string;
    setSpecificView: Function;
};

export type TagType = {
    id: string;
    name: string;
}

export function FormResponseView(props: FormResponseViewProps) {
    const [isResolved, setIsResolved] = useState(true);
    const [tags, setTags] = useState<TagType[]>([{id: "1", name: "test"}]);
    const [mailto, setMailto] = useState("mailto:");
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

    const getTags = async(id: string) => {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + id + Endpoints.ResponsesTags, {
            method: "GET",
            headers: new Headers({
                "Authorization": authToken
            })
        });
        if (response.status >= 300) {
            console.log("Error retreiving response tags");
            return;
        }
        let tagsList: TagType[] = [];
        const tags = await response.json();
        tags.forEach(function(tag: any) {
            tagsList.push({id: tag.id, name: tag.name});
        });
        setTags(tagsList);
    }

    async function removeTag(responseId: string) {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + responseId + Endpoints.ResponsesTags, {

        });
    }

    async function addTag(responseId: string) {
        var authToken = localStorage.getItem("Authorization") || "";
        const response = await fetch(Endpoints.Testbase + Endpoints.Responses + "/" + responseId + Endpoints.ResponsesTags, {

        });
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

    let tagsList:any[] = [];
    tags.forEach(function(tag) {
        tagsList.push(<Tag tagId={tag.id} name={tag.name}></Tag>)
    });

    return(
        <div>
            <button className="exit-button" onClick={() => props.setSpecificView()}><img src="./assets/icons/back-arrow.png"></img></button>
            <Header title={props.title}></Header>
            <SubHeaderLine title={props.subject}></SubHeaderLine>
            <div className="tags-container">
                    {tagsList}
                    <TagAdd></TagAdd>
            </div>
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