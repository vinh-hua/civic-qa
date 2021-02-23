export type HeaderProps = {
    title: string;
}

export function Header(props: HeaderProps) {
    return(
        <div>
            <h1 className="sub-dash-title">{props.title}</h1>
        </div>
    );
}