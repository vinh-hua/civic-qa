import { Dispatch, SetStateAction } from 'react';
import { SubDashboardCard } from './SubDashboardCard';
import { SubHeaderLine } from './SubHeaderLine';
import './SubDashboard.css';
import { EmailTemplateCard } from './EmailTemplateCard';

export type SubDashboardData = {
    name: string;
    value: any;
}

export type SubDashboardProps = {
    title: string;
    data: Array<SubDashboardData>;
    setData: Dispatch<SetStateAction<SubDashboardData[]>>;
    emailTemplates: boolean;
    hasRespondOption: boolean;
};

export function SubDashboard(props: SubDashboardProps) {
    let cards:any[] = [];
    if (props.emailTemplates) {
        props.data.forEach(d => cards.push(<EmailTemplateCard name={d.name} value={d.value}></EmailTemplateCard>))
    } else {
        props.data.forEach(d => cards.push(<SubDashboardCard name={d.name} value={d.value} setData={props.setData} hasRespondOption={props.hasRespondOption}></SubDashboardCard>));
    }

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