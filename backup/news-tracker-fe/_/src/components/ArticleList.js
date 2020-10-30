import React from "react";
import Article from "./Article";

function ArticleList(props) {
  return (
    <div>
    {props.news.map(c => <Article key={c.author} item={c}  />)}
    xx
   </div>  
  );
}

export default ArticleList;