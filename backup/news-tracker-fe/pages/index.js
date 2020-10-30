import Head from 'next/head'
import styles from '../styles/Home.module.css'
import useSwr from 'swr'
import ArticleList from '../components/ArticleList';

const fetcher = (url) => fetch(url).then((res) => res.json())

export default function Home() {
  const { data, error } = useSwr('/api/news', fetcher);
  return (
    <div className="App" >
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
