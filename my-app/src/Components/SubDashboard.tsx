import { SubDashboardCard } from './SubDashboardCard';
import './SubDashboard.css';

export type SubDashboardData = {
    name: string;
    value: number;
}

export type SubDashboardProps = {
    title: string;
    data: Array<SubDashboardData>;
};

export function SubDashboard(props: SubDashboardProps) {
    let cards:any[] = [];
    props.data.forEach(d => cards.push(<SubDashboardCard name={d.name} value={d.value}></SubDashboardCard>));

    return (
        <div>
            <div>
                <h2 className="sub-dash-sub-title">{props.title}</h2>
                <hr className="sub-dash-line" />
                <div className="sub-dash-cards">
                    {cards}
                </div>
            </div>
        </div>
    );
}