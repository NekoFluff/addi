// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetInArtHandlerFunc turns a function with the right signature into a get in art handler
type GetInArtHandlerFunc func(GetInArtParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInArtHandlerFunc) Handle(params GetInArtParams) middleware.Responder {
	return fn(params)
}

// GetInArtHandler interface for that can handle valid get in art params
type GetInArtHandler interface {
	Handle(GetInArtParams) middleware.Responder
}

// NewGetInArt creates a new http.Handler for the get in art operation
func NewGetInArt(ctx *middleware.Context, handler GetInArtHandler) *GetInArt {
	return &GetInArt{Context: ctx, Handler: handler}
}

/* GetInArt swagger:route GET /inArt getInArt

Get inART

Get inART

*/
type GetInArt struct {
	Context *middleware.Context
	Handler GetInArtHandler
}

func (o *GetInArt) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetInArtParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}