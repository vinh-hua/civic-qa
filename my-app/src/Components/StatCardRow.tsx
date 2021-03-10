import { StatCard, StatCardProps } from './StatCard';
import './StatCardRow.css';

export type StatCardRowProps = {
    cards: Array<StatCardProps>;
    spaceEven: boolean;
}

export function StatCardRow(props: StatCardRowProps) {
    let statCards:any[] = [];
    props.cards.forEach(card => statCards.push(<StatCard title={card.title} stat={card.stat}></StatCard>));

    return(
        props.spaceEven ? <div className="stat-cards-even">
            {statCards}
        </div> :
        <div className="stat-cards-between">
            {statCards}
        </div>
    );
}