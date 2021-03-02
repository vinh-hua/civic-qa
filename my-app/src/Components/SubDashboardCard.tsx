import { Dispatch, SetStateAction } from 'react';
import { SubDashboardData } from './SubDashboard';
import './SubDashboardCard.css';

export type SubDashboardCardProps = {
    name: string;
    value: number;
    setData: Dispatch<SetStateAction<SubDashboardData[]>>;
    viewButton: boolean;
    hasRespondOption: boolean;
}

// using test data for now, integrate with API data here
function getSubDashboardData(): Array<SubDashboardData> {
    var data = [];
    data.push({name: "test", value: 123});
    data.push({name: "test", value: 119});
    data.push({name: "test", value: 77});
    data.push({name: "test", value: 62});
    data.push({name: "test", value: 36});
    data.push({name: "test", value: 52});
    data.push({name: "test", value: 52});
    data.push({name: "test", value: 52});
    data.push({name: "test", value: 52});
    return data as Array<SubDashboardData>;
}

export function SubDashboardCard(props: SubDashboardCardProps) {
    let test_data = getSubDashboardData();

    const buttonName = props.viewButton ? "View" : "Respond";

    return (
        <div className="sub-dash-card">
            <p className="sub-dash-card-name">{props.name}</p>
            <button className="sub-dash-card-btn" onClick={() => props.setData(test_data)}>{buttonName}</button>
            <p className="sub-dash-card-value">{props.value}</p>
            {props.hasRespondOption ? <button className="write-btn"><img src="./assets/icons/write.png"></img></button> : null}
        </div>
    );
}