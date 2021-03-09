import { SubDashboardData } from '../Components/SubDashboard';
import './SubDashboardCard.css';

export type SubDashboardCardProps = {
    data: SubDashboardData;
    changeViewFunc: Function;
}

export function SubDashboardCard(props: SubDashboardCardProps) {
    return (
        <div className="sub-dash-card">
            <p className="sub-dash-card-name">{props.data.name}</p>
            <button className="sub-dash-card-btn" onClick={() => props.changeViewFunc(props.data)}>View</button>
            <p className="sub-dash-card-value">{props.data.value}</p>
        </div>
    );
}