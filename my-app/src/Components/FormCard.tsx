import './FormCard.css';

export type FormCardProps = {
    id: string;
    name: string;
    getForm: Function;
}

export function FormCard(props: FormCardProps) {

    return(
        <div className="form-card">
            <h1 className="form-card-name">{props.name}</h1>
            <button className="form-card-btn" onClick={() => props.getForm(props.id)}>Copy iFrame embed link</button>
        </div>
    );
}