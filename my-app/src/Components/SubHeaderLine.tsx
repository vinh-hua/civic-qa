export type SubHeaderLineProps = {
    title: string;
}

export function SubHeaderLine(props: SubHeaderLineProps) {
    return(
        <div>
            <h2 className="sub-dash-sub-title">{props.title}</h2>
            <hr className="sub-dash-line" />
        </div>
    );
}