import { Dispatch, SetStateAction } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from './SubHeaderLine';
import "./SpecificView.css";

export type SpecificViewProps = {
    title: string;
    subHeaderNumber: number;
    subject: string;
    body: string;
    setSpecificView: Dispatch<SetStateAction<boolean>>;
};

export function SpecificView(props: SpecificViewProps) {
    return(
        <div>
            <Header title={props.title}></Header>
            <SubHeaderLine title={props.subject}></SubHeaderLine>
            <button className="exit-button" onClick={() => props.setSpecificView(false)}>X</button>
            <div className="form-response">
                <p className="form-response-body">{props.body}</p>
                <textarea className="form-response-message"></textarea>
                <button className="send-btn" placeholder="Message">Send</button>
            </div>
        </div>
    );
}