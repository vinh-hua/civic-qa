import { StatCard, StatCardProps } from './StatCard';
import './StatCardRow.css';

export type StatCardRowProps = {
    cards: Array<StatCardProps>;
}

export function StatCardRow(props: StatCardRowProps) {
    let statCards:any[] = [];
    props.cards.forEach(card => statCards.push(<StatCard title={card.title} stat={card.stat}></StatCard>));

    return(
        <div className="stat-cards">
            {statCards}
        </div>
    );
}