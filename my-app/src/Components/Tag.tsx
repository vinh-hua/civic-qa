import './Tag.css';

export type TagProps = {
    name: string;
}

export function Tag(props: TagProps) {
    return(
        <div className="tag">
            <p className="tag-name">{"#" + props.name}</p>
            <button className="tag-remove-btn"><img src="./assets/icons/remove.png"></img></button>
        </div>
    );
}