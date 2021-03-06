import { Dispatch, SetStateAction } from 'react';
import './SubDashboardCard.css';

export type SubDashboardCardProps = {
    name: string;
    value: number;
    setSpecificView: Function;
    viewButton: boolean;
}

export function SubDashboardCard(props: SubDashboardCardProps) {
    const buttonName = props.viewButton ? "View" : "Respond";

    return (
        <div className="sub-dash-card">
            <p className="sub-dash-card-name">{props.name}</p>
            <button className="sub-dash-card-btn" onClick={() => props.setSpecificView(props.name)}>{buttonName}</button>
            <p className="sub-dash-card-value">{props.value}</p>
        </div>
    );
}