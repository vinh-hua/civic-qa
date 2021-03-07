import { Dispatch, SetStateAction } from 'react';
import { Header } from './Header';
import { SubHeaderLine } from './SubHeaderLine';
import "./FormResponseView.css";

export type FormResponseViewProps = {
    title: string;
    subHeaderNumber: number;
    subject: string;
    body: string;
    setSpecificView: Dispatch<SetStateAction<boolean>>;
};

export function FormResponseView(props: FormResponseViewProps) {
    return(
        <div>
            <button className="exit-button" onClick={() => props.setSpecificView(false)}><img src="./assets/icons/back-arrow.png"></img></button>
            <Header title={props.title}></Header>
            <SubHeaderLine title={props.subject}></SubHeaderLine>
            <div className="form-response">
                <p className="form-response-body">{props.body}</p>
                <textarea className="form-response-message"></textarea>
                <div className="resolved-send-container">
                    <label className="resolved-label" >
                        <input id="resolved-check-box" className="resolved-check-box" type="checkbox"></input>
                        Resolved
                    </label>
                    <button className="send-btn" placeholder="Message">Send</button>
                </div>
            </div>
        </div>
    );
}