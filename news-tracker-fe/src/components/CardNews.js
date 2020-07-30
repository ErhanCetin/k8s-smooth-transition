import React from 'react';
import { MDBBtn, MDBCard, MDBCardBody, MDBCardImage, MDBCardTitle, MDBCardText, MDBCol } from 'mdbreact';

const CardNews = (props) => {
  return (
    <MDBCol>
      <MDBCard style={{ width: "22rem" }}>
        <MDBCardImage className="img-fluid" src={props.item.urlToImage} waves />
        <MDBCardBody>
          <MDBCardTitle>{props.item.title}</MDBCardTitle>
          <MDBCardText>{props.item.description}
          </MDBCardText>
          <MDBBtn href={props.item.url}>{props.item.source_name}</MDBBtn>
        </MDBCardBody>
      </MDBCard>
    </MDBCol>
  )
}
export default CardNews;
