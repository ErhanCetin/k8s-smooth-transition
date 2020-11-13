import Head from 'next/head'
import styles from '../styles/Home.module.css'
import useSwr from 'swr'
import ArticleList from '../components/ArticleList';
import 'semantic-ui-css/semantic.min.css'




const fetcher = (url) => fetch(url).then((res) => res.json())

export default function Home() {
  const { data, error } = useSwr('/api/news', fetcher);
  return (
    <div>

<div class="ui horizontal divider"/>
<span>
  <i class="linkedin icon"/>
    <a href="https://www.linkedin.com/in/erhancetin/">erhancetin</a>
  <br/>
  <i class="github icon"/>
  <a href="https://github.com/ErhanCetin/k8s-smooth-transition">k8s-smooth-transition</a>
</span>

<h2 class="ui center aligned icon header">
  <i class="paw icon"></i>
  </h2>
        <h3 class="ui horizontal divider header">
          <i class="tag icon"/>Welcome News Tracker Application
        </h3>
      <ArticleList
        news={(data || []).map(c => ({
          id : c.id,
          author: c.author,
          title: c.title,
          description: c.description,
          url: c.url,
          urlToImage: c.urltoimage,
          publishedAt: c.publishedat,
          content: c.content ,
          source_name: c.source.name
        }))}
      />
   </div>
  )
}
