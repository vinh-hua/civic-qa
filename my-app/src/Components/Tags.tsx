import { useEffect, useRef, useState } from 'react';
import { Tag } from './Tag';
import './Tag.css';

export type TagProps = {
    values: string[];
    addTag: Function;
    removeTag: Function;
}

export function Tags(props: TagProps) {
    const wrapperRef = useRef<HTMLInputElement>(null);
    const [inputText, setInputText] = useState("");
    const [inputShow, setInputShow] = useState(false);

    function addNewTag() {
        if (inputText.length > 0) {
            props.addTag(inputText);
        }
        setInputShow(false);
        setInputText("");
    }

    function enterNewTag(e: any) {
        e.preventDefault();
        addNewTag();
    }  

    function handleClickOutside(e: any) {
        if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
            setInputShow(false);
        }
    }

    useEffect(() => {
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);

    let tagList: any[] = [];
    props.values.forEach(function(value) {
        tagList.push(<Tag value={value} remove={props.removeTag}></Tag>)
    });

    return(
        <div className="tags-container">
            {tagList}
            {inputShow ? <div>
                    <form onSubmit={e => enterNewTag(e)}>
                        <input ref={wrapperRef} className="tag-add-input" type="text" placeholder={"#"} onChange={e => setInputText(e.target.value)}></input>
                    </form>
                </div>
                : <button className="tag-add-btn" onClick={() => setInputShow(true)}><img src="./assets/icons/add.png"></img></button>}
        </div>
    );
}