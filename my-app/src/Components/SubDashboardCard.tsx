import './SubDashboardCard.css';

export type SubDashboardCardProps = {
    name: string;
    value: number;
    changeViewFunc: Function;
}

export function SubDashboardCard(props: SubDashboardCardProps) {
    return (
        <div className="sub-dash-card">
            <p className="sub-dash-card-name">{props.name}</p>
            <button className="sub-dash-card-btn" onClick={() => props.changeViewFunc()}>View</button>
            <p className="sub-dash-card-value">{props.value}</p>
        </div>
    );
}