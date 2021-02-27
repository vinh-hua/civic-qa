import './EmailTemplateCard.css';

export type EmailTemplateCardProps = {
    name: string;
    value: number;
}

export function EmailTemplateCard(props: EmailTemplateCardProps) {
    return (
        <div className="email-tmp-card">
            <p className="email-tmp-card-name">{props.name}</p>
            <button className="email-tmp-card-btn">View</button>
            <p className="email-tmp-card-value">{props.value}</p>
        </div>
    );
}