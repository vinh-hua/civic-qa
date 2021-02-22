export type SubHeaderProps = {
    title: string;
}

export function SubHeader(props: SubHeaderProps) {
    return(
        <div>
            <h1 className="sub-dash-title">{props.title}</h1>
        </div>
    );
}