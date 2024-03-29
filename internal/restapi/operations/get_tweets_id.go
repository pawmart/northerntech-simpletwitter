// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTweetsIDHandlerFunc turns a function with the right signature into a get tweets ID handler
type GetTweetsIDHandlerFunc func(GetTweetsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTweetsIDHandlerFunc) Handle(params GetTweetsIDParams) middleware.Responder {
	return fn(params)
}

// GetTweetsIDHandler interface for that can handle valid get tweets ID params
type GetTweetsIDHandler interface {
	Handle(GetTweetsIDParams) middleware.Responder
}

// NewGetTweetsID creates a new http.Handler for the get tweets ID operation
func NewGetTweetsID(ctx *middleware.Context, handler GetTweetsIDHandler) *GetTweetsID {
	return &GetTweetsID{Context: ctx, Handler: handler}
}

/*GetTweetsID swagger:route GET /tweets/{id} getTweetsId

Fetch tweet

*/
type GetTweetsID struct {
	Context *middleware.Context
	Handler GetTweetsIDHandler
}

func (o *GetTweetsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTweetsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
