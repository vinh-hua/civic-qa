import { SubDashboardCard } from './SubDashboardCard';
import './SubDashboard.css';

export type SubDashboardData = {
    name: string;
    value: number;
}

export type SubDashboardProps = {
    title: string;
    subTitle: string;
    data: Array<SubDashboardData>;
};

export function SubDashboard(props: SubDashboardProps) {
    let cards:any[] = [];
    props.data.forEach(d => cards.push(<SubDashboardCard name={d.name} value={d.value}></SubDashboardCard>));

    return (
        <div>
            <h1 className="sub-dash-title">{props.title}</h1>
            <div>
                <h2 className="sub-dash-sub-title">{props.subTitle}</h2>
                <hr className="sub-dash-line" />
                <div className="sub-dash-cards">
                    {cards}
                </div>
            </div>


        </div>
    );
}