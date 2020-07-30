import React from "react";
// eslint-disable-next-line
import CardNews from "./CardNews";
import BlogNews from "./BlogNews";


function Article(props) {
    return (

        <div >
            
            <BlogNews item={props.item} />
        </div >

    );
}
export default Article;