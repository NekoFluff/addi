// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDspHandlerFunc turns a function with the right signature into a get dsp handler
type GetDspHandlerFunc func(GetDspParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDspHandlerFunc) Handle(params GetDspParams) middleware.Responder {
	return fn(params)
}

// GetDspHandler interface for that can handle valid get dsp params
type GetDspHandler interface {
	Handle(GetDspParams) middleware.Responder
}

// NewGetDsp creates a new http.Handler for the get dsp operation
func NewGetDsp(ctx *middleware.Context, handler GetDspHandler) *GetDsp {
	return &GetDsp{Context: ctx, Handler: handler}
}

/* GetDsp swagger:route GET /dsp getDsp

Get the optimal recipe

*/
type GetDsp struct {
	Context *middleware.Context
	Handler GetDspHandler
}

func (o *GetDsp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDspParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}