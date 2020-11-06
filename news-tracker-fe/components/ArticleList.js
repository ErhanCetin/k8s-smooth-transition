import React from "react";
import Article from "./Article";

function ArticleList(props) {
  return (
    <div>
     {props.news.map(c => <Article key={c.id} item={c} />)}
    -- my place holder --
   </div>  
  );
}
export default ArticleList;