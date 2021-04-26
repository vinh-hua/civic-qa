import { SubDashboardCard } from './SubDashboardCard';
import { SubHeaderLine } from './SubHeaderLine';
import './SubDashboard.css';
import { EmailTemplateCard } from './EmailTemplateCard';

export type InquiryData = {
    id?: string;
    email?: string;
    name: string;
    value: any;
    body?: string;
}

export type SubDashboardProps = {
    title: string;
    data: Array<InquiryData>;
    changeViewFunc: Function;
    fullPageView: boolean;
    subHeaderValue?: any;
};

export function SubDashboard(props: SubDashboardProps) {
    let cards: any[] = [];
    props.data.forEach(d => cards.push(<SubDashboardCard data={d} changeViewFunc={props.changeViewFunc}></SubDashboardCard>));

    return (
        <div>
            <div>
                <SubHeaderLine title={props.title} subHeaderValue={props.subHeaderValue ? props.subHeaderValue : undefined}></SubHeaderLine>
                {props.fullPageView ? 
                <div className="sub-dash-cards-long">
                    {cards}
                </div> :
                <div className="sub-dash-cards-short">
                    {cards}
                </div>
                }

            </div>
        </div>
    );
}