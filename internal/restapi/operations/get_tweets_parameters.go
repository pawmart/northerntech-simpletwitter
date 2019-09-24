// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTweetsParams creates a new GetTweetsParams object
// no default values defined in spec.
func NewGetTweetsParams() GetTweetsParams {

	return GetTweetsParams{}
}

// GetTweetsParams contains all the bound params for the get tweets operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTweets
type GetTweetsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Count
	  In: query
	*/
	Count *int64
	/*Filter by tag
	  In: query
	*/
	FilterTag []string
	/*Year
	  In: query
	*/
	Year *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTweetsParams() beforehand.
func (o *GetTweetsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilterTag, qhkFilterTag, _ := qs.GetOK("filter[tag]")
	if err := o.bindFilterTag(qFilterTag, qhkFilterTag, route.Formats); err != nil {
		res = append(res, err)
	}

	qYear, qhkYear, _ := qs.GetOK("year")
	if err := o.bindYear(qYear, qhkYear, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCount binds and validates parameter Count from query.
func (o *GetTweetsParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("count", "query", "int64", raw)
	}
	o.Count = &value

	return nil
}

// bindFilterTag binds and validates array parameter FilterTag from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetTweetsParams) bindFilterTag(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvFilterTag string
	if len(rawData) > 0 {
		qvFilterTag = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	filterTagIC := swag.SplitByFormat(qvFilterTag, "")
	if len(filterTagIC) == 0 {
		return nil
	}

	var filterTagIR []string
	for _, filterTagIV := range filterTagIC {
		filterTagI := filterTagIV

		filterTagIR = append(filterTagIR, filterTagI)
	}

	o.FilterTag = filterTagIR

	return nil
}

// bindYear binds and validates parameter Year from query.
func (o *GetTweetsParams) bindYear(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("year", "query", "int64", raw)
	}
	o.Year = &value

	return nil
}