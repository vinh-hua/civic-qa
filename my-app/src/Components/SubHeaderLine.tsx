export type SubHeaderLineProps = {
    title: string;
    subHeaderNumber?: number;
}

export function SubHeaderLine(props: SubHeaderLineProps) {
    return(
        <div>
            <div className="sub-dash-sub-header">
                <h2 className="sub-dash-sub-title">{props.title}</h2>
                {props.subHeaderNumber ? <h1 className="sub-dash-sub-number">{props.subHeaderNumber}</h1> : null}
            </div>
            <hr className="sub-dash-line" />
        </div>
    );
}