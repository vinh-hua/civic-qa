import './FormCard.css';

export type FormCardProps = {
    id: string;
    name: string;
    getForm: Function;
}

export function FormCard(props: FormCardProps) {

    return(
        <div>
            <h1>{props.name}</h1>
            <button onClick={() => props.getForm(props.id)}>Get Embed Form Link</button>
        </div>
    );
}