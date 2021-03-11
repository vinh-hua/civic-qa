import './Tag.css';

export type TagProps = {
    value: string;
}

export function Tag(props: TagProps) {
    return(
        <div className="tag">
            <p className="tag-value">{"#" + props.value}</p>
            <button className="tag-remove-btn"><img src="./assets/icons/remove.png"></img></button>
        </div>
    );
}