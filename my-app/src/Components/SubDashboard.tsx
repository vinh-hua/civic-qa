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
    setSpecificView: Function;
    emailTemplates: boolean;
    fullPageView: boolean;
    viewButton: boolean;
    subHeaderNumber?: number;
};

export function SubDashboard(props: SubDashboardProps) {
    let cards:any[] = [];
    if (props.emailTemplates) {
        props.data.forEach(d => cards.push(<EmailTemplateCard name={d.name} value={d.value}></EmailTemplateCard>))
    } else {
        props.data.forEach(d => cards.push(<SubDashboardCard name={d.name} value={d.value} setSpecificView={props.setSpecificView} viewButton={props.viewButton}></SubDashboardCard>));
    }

    return (
        <div>
            <div>
                <SubHeaderLine title={props.title} subHeaderNumber={props.subHeaderNumber ? props.subHeaderNumber : undefined}></SubHeaderLine>
                {props.fullPageView ? 
                <div className="sub-dash-cards-700">
                    {cards}
                </div> :
                <div className="sub-dash-cards-400">
                    {cards}
                </div>
                }

            </div>
        </div>
    );
}