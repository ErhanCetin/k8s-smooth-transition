import React from "react";
// eslint-disable-next-line
import CardNews from "./CardNews";
// eslint-disable-next-line
// import BlogNews from "./BlogNews";


function Article(props) {
    return (
        <div >
            <CardNews item={props.item} />
        </div >
    );
}
export default Article;