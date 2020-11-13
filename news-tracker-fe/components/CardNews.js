import React from 'react';
import { Card, Icon, Image ,Item} from 'semantic-ui-react'

const CardNews = ({ item }) => {
  return (

    <Card>
    <Image src={item.urlToImage} wrapped ui={false} />
    <Card.Content>
      <Card.Header>{item.title}</Card.Header>
      <Card.Meta><span className='date'><Icon name="calendar alternate icon"/>{item.publishedAt}</span></Card.Meta>
      <Card.Description>{item.description}</Card.Description>
    </Card.Content>
    <Card.Content extra>
     <span><a href={item.url} ><Icon name='list alternate outline' />Source : {item.source_name}</a></span>
      <br/>
      <span><Icon name='user circle icon' />Author: {item.author}</span>
    </Card.Content>
  </Card>  
  )
}
export default CardNews;

