// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pawmart/northerntech-simpletwitter/internal/models"
)

// PatchTweetsOKCode is the HTTP code returned for type PatchTweetsOK
const PatchTweetsOKCode int = 200

/*PatchTweetsOK Tweet update response

swagger:response patchTweetsOK
*/
type PatchTweetsOK struct {
}

// NewPatchTweetsOK creates PatchTweetsOK with default headers values
func NewPatchTweetsOK() *PatchTweetsOK {

	return &PatchTweetsOK{}
}

// WriteResponse to the client
func (o *PatchTweetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchTweetsBadRequestCode is the HTTP code returned for type PatchTweetsBadRequest
const PatchTweetsBadRequestCode int = 400

/*PatchTweetsBadRequest Tweet update error

swagger:response patchTweetsBadRequest
*/
type PatchTweetsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPatchTweetsBadRequest creates PatchTweetsBadRequest with default headers values
func NewPatchTweetsBadRequest() *PatchTweetsBadRequest {

	return &PatchTweetsBadRequest{}
}

// WithPayload adds the payload to the patch tweets bad request response
func (o *PatchTweetsBadRequest) WithPayload(payload *models.APIError) *PatchTweetsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tweets bad request response
func (o *PatchTweetsBadRequest) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTweetsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTweetsNotFoundCode is the HTTP code returned for type PatchTweetsNotFound
const PatchTweetsNotFoundCode int = 404

/*PatchTweetsNotFound Tweet update not found

swagger:response patchTweetsNotFound
*/
type PatchTweetsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPatchTweetsNotFound creates PatchTweetsNotFound with default headers values
func NewPatchTweetsNotFound() *PatchTweetsNotFound {

	return &PatchTweetsNotFound{}
}

// WithPayload adds the payload to the patch tweets not found response
func (o *PatchTweetsNotFound) WithPayload(payload *models.APIError) *PatchTweetsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tweets not found response
func (o *PatchTweetsNotFound) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTweetsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchTweetsInternalServerErrorCode is the HTTP code returned for type PatchTweetsInternalServerError
const PatchTweetsInternalServerErrorCode int = 500

/*PatchTweetsInternalServerError Fatal error

swagger:response patchTweetsInternalServerError
*/
type PatchTweetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewPatchTweetsInternalServerError creates PatchTweetsInternalServerError with default headers values
func NewPatchTweetsInternalServerError() *PatchTweetsInternalServerError {

	return &PatchTweetsInternalServerError{}
}

// WithPayload adds the payload to the patch tweets internal server error response
func (o *PatchTweetsInternalServerError) WithPayload(payload *models.APIError) *PatchTweetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch tweets internal server error response
func (o *PatchTweetsInternalServerError) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchTweetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
