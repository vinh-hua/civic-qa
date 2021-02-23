import { SubDashboardCard } from './SubDashboardCard';
import { SubHeaderLine } from './SubHeaderLine';
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
                <SubHeaderLine title={props.title}></SubHeaderLine>
                <div className="sub-dash-cards">
                    {cards}
                </div>
            </div>
        </div>
    );
}