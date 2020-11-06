
import React from "react";
// eslint-disable-next-line
import { MDBRow, MDBCol, MDBCard, MDBCardBody, MDBMask, MDBIcon, MDBView, MDBBtn } from "mdbreact";

const BlogNews = (props) => {
  return (
    <MDBCard className="my-5 px-5 pb-5">
      <MDBCardBody>
       <MDBRow>
          <MDBCol lg="5" xl="4">
            <MDBView hover className="rounded z-depth-1-half mb-lg-0 mb-4">
              <img
                className="rounded float-left"
                src={props.item.urlToImage}
                alt=""
                size="200px"
                            
              />
              <a href="#!">
                <MDBMask overlay="white-slight" />
              </a>
            </MDBView>
          </MDBCol>
          <MDBCol lg="7" xl="8">
            <h3 className="font-weight-bold mb-3 p-0">
              <strong>{props.item.title}</strong>
            </h3>
            <p className="dark-grey-text">
                {props.item.description}
            </p>
            <p>
              by <a href="#!" className="font-weight-bold">{props.item.author}</a>, {props.item.publishedat}
            </p>
            <MDBBtn color="primary" size="md" href={props.item.url}>
              Read More
            </MDBBtn>
          </MDBCol>
        </MDBRow>
        <hr className="my-5" />
      </MDBCardBody>
    </MDBCard>
  );
}

export default BlogNews;
