import { Dispatch, SetStateAction } from 'react';
import { Header } from '../Components/Header';
import { SubHeaderLine } from './SubHeaderLine';

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
            <SubHeaderLine title={props.title} subHeaderNumber={props.subHeaderNumber}></SubHeaderLine>
            <div>
                <button className="exit-button" onClick={() => props.setSpecificView(false)}></button>
                <p>{props.subject}</p>
                <p>{props.body}</p>
            </div>
        </div>
    );
}