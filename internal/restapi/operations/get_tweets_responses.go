// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pawmart/northerntech-simpletwitter/internal/models"
)

// GetTweetsOKCode is the HTTP code returned for type GetTweetsOK
const GetTweetsOKCode int = 200

/*GetTweetsOK List of tweet details

swagger:response getTweetsOK
*/
type GetTweetsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TweetDetailsListResponse `json:"body,omitempty"`
}

// NewGetTweetsOK creates GetTweetsOK with default headers values
func NewGetTweetsOK() *GetTweetsOK {

	return &GetTweetsOK{}
}

// WithPayload adds the payload to the get tweets o k response
func (o *GetTweetsOK) WithPayload(payload *models.TweetDetailsListResponse) *GetTweetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tweets o k response
func (o *GetTweetsOK) SetPayload(payload *models.TweetDetailsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTweetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTweetsInternalServerErrorCode is the HTTP code returned for type GetTweetsInternalServerError
const GetTweetsInternalServerErrorCode int = 500

/*GetTweetsInternalServerError Fatal error

swagger:response getTweetsInternalServerError
*/
type GetTweetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewGetTweetsInternalServerError creates GetTweetsInternalServerError with default headers values
func NewGetTweetsInternalServerError() *GetTweetsInternalServerError {

	return &GetTweetsInternalServerError{}
}

// WithPayload adds the payload to the get tweets internal server error response
func (o *GetTweetsInternalServerError) WithPayload(payload *models.APIError) *GetTweetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tweets internal server error response
func (o *GetTweetsInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTweetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}