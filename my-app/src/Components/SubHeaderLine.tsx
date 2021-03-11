export type SubHeaderLineProps = {
    title: string;
    subHeaderValue?: any;
}

export function SubHeaderLine(props: SubHeaderLineProps) {
    return(
        <div>
            <div className="sub-dash-sub-header">
                <h2 className="sub-dash-sub-title">{props.title}</h2>
                {props.subHeaderValue ? <h1 className="sub-dash-sub-number">{props.subHeaderValue}</h1> : null}
                <hr className="sub-dash-line" />
            </div>
        </div>
    );
}