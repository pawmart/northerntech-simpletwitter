// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pawmart/northerntech-simpletwitter/internal/models"
)

// PostTweetsCreatedCode is the HTTP code returned for type PostTweetsCreated
const PostTweetsCreatedCode int = 201

/*PostTweetsCreated Tweet creation response

swagger:response postTweetsCreated
*/
type PostTweetsCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TweetCreationResponse `json:"body,omitempty"`
}

// NewPostTweetsCreated creates PostTweetsCreated with default headers values
func NewPostTweetsCreated() *PostTweetsCreated {

	return &PostTweetsCreated{}
}

// WithPayload adds the payload to the post tweets created response
func (o *PostTweetsCreated) WithPayload(payload *models.TweetCreationResponse) *PostTweetsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tweets created response
func (o *PostTweetsCreated) SetPayload(payload *models.TweetCreationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTweetsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTweetsBadRequestCode is the HTTP code returned for type PostTweetsBadRequest
const PostTweetsBadRequestCode int = 400

/*PostTweetsBadRequest Tweet creation error

swagger:response postTweetsBadRequest
*/
type PostTweetsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostTweetsBadRequest creates PostTweetsBadRequest with default headers values
func NewPostTweetsBadRequest() *PostTweetsBadRequest {

	return &PostTweetsBadRequest{}
}

// WithPayload adds the payload to the post tweets bad request response
func (o *PostTweetsBadRequest) WithPayload(payload *models.APIError) *PostTweetsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tweets bad request response
func (o *PostTweetsBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTweetsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTweetsInternalServerErrorCode is the HTTP code returned for type PostTweetsInternalServerError
const PostTweetsInternalServerErrorCode int = 500

/*PostTweetsInternalServerError Fatal error

swagger:response postTweetsInternalServerError
*/
type PostTweetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPostTweetsInternalServerError creates PostTweetsInternalServerError with default headers values
func NewPostTweetsInternalServerError() *PostTweetsInternalServerError {

	return &PostTweetsInternalServerError{}
}

// WithPayload adds the payload to the post tweets internal server error response
func (o *PostTweetsInternalServerError) WithPayload(payload *models.APIError) *PostTweetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tweets internal server error response
func (o *PostTweetsInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTweetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}