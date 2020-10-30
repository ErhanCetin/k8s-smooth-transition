import React from 'react';
import { MDBBtn, MDBCard, MDBCardBody, MDBCardImage, MDBCardTitle, MDBCardText, MDBCol } from 'mdbreact';

const CardNews = ({ item }) => {
  return (
    <MDBCol>
      <MDBCard style={{ width: "22rem" }}>
        <MDBCardImage className="img-fluid" src={item.urlToImage} waves />
        <MDBCardBody>
          <MDBCardTitle>{item.title}</MDBCardTitle>
          <MDBCardText>{item.description}
          </MDBCardText>
          <MDBBtn href={item.url}>{item.source_name}</MDBBtn>
        </MDBCardBody>
      </MDBCard>
    </MDBCol>
  )
}
export default CardNews;
