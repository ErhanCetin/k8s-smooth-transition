import React from "react";
import CardNews from "./CardNews";


function Article(props) {
    return (
        <div>
            <CardNews item={props.item} />
        </div>
    );
}
export default Article;