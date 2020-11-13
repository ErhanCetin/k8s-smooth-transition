import React from "react";
import Article from "./Article";
import CardNews from "./CardNews";


function ArticleList(props) {
  return (
    <div class="ui four cards">
     {props.news.map(c => <CardNews item={c}/>)}
     -- my place holder --
   </div>  
  );
}
export default ArticleList;