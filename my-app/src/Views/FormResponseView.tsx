import { Dispatch, SetStateAction, useState } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from '../Components/SubHeaderLine';
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
    const [isResolved, setIsResolved] = useState(false);

    const resolveResponse = async(id: string, isResolved: boolean) => {
        var authToken = localStorage.getItem("Authorization") || "";
        var patchActive = JSON.stringify({isActive: !isResolved});
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

    return(
        <div>
            <button className="exit-button" onClick={() => props.setSpecificView()}><img src="./assets/icons/back-arrow.png"></img></button>
            <Header title={props.title}></Header>
            <SubHeaderLine title={props.subject}></SubHeaderLine>
            <div className="form-response">
                <p className="form-response-body">{props.body}</p>
                <textarea className="form-response-message"></textarea>
                <div className="resolved-send-container">
                    <label className="resolved-label" >
                        <input id="resolved-check-box" className="resolved-check-box" type="checkbox" onClick={() => clickCheckbox()}></input>
                        Resolved
                    </label>
                    <button className="send-btn" placeholder="Message">Send</button>
                </div>
            </div>
        </div>
    );
}