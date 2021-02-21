import './StatCard.css';

export type StatCardProps = {
    title: string;
    stat: number;
};

export function StatCard(props: StatCardProps) {
    return (
        <button className="stat-card">
            <h1 className="stat-card-title">{props.title}</h1>
            <p className="stat-card-data">{props.stat}</p>
        </button>
    );
}