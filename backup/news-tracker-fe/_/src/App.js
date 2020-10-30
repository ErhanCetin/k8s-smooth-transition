import React, { Component } from 'react';
import axios from "axios";
import ArticleList from "./components/ArticleList";
import Config from './Config';

class App extends Component {

    // default State object
    state = {
        news: []
    };

    componentDidMount() {

        // var newsApiUrl = 'http://'+Config.newsapiHostName+':'+Config.newsApiHostPort+'/news/getAll';
        // console.log('xxxxxxx', newsApiUrl );
        // console.log(Config.newsapiHostName + ' ::: '+ Config.newsApiHostPort);
        axios
            .get('/api/news'
            ).then(response => {

                // create an array of contacts only with relevant data
                const newArticles = response.data.map(c => {
                    return {
                        id : c.id,
                        author: c.author,
                        title: c.title,
                        description: c.description,
                        url: c.url,
                        urlToImage: c.urltoimage,
                        publishedAt: c.publishedat,
                        content: c.content ,
                        source_name: c.source.name
                    }});
                    
                    // create a new "State" object without mutating 
                    // the original State object. 
                    const newState = Object.assign({}, this.state, {
                        news: newArticles
                    });

                    // store the new state object in the component's state
                    this.setState(newState);
                })
                    .catch(error => console.log(error));
           
    }

    render() {
        return (
            <div className="App" >
                <ArticleList news={this.state.news} />
            </div>
        );
    }
}

export default App