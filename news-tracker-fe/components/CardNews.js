import React from 'react';
import { Card, CardGroup, Icon, Image ,Button, Grid, Segment,Item} from 'semantic-ui-react'

const CardNews = ({ item }) => {
  return (

    <Card>
    <Image src={item.urlToImage} wrapped ui={false} />
    <Card.Content>
      <Card.Header>{item.title}</Card.Header>
      <Card.Meta>
        <span className='date'>Joined in 2015</span>
      </Card.Meta>
      <Card.Description>
      {item.description}
      </Card.Description>
    </Card.Content>
    <Card.Content extra>
      <a>
        <Icon name='user' />
        {item.source_name}
      </a>
    </Card.Content>
  </Card>  
  )
}
export default CardNews;

