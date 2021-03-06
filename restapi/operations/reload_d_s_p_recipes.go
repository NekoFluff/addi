// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReloadDSPRecipesHandlerFunc turns a function with the right signature into a reload d s p recipes handler
type ReloadDSPRecipesHandlerFunc func(ReloadDSPRecipesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReloadDSPRecipesHandlerFunc) Handle(params ReloadDSPRecipesParams) middleware.Responder {
	return fn(params)
}

// ReloadDSPRecipesHandler interface for that can handle valid reload d s p recipes params
type ReloadDSPRecipesHandler interface {
	Handle(ReloadDSPRecipesParams) middleware.Responder
}

// NewReloadDSPRecipes creates a new http.Handler for the reload d s p recipes operation
func NewReloadDSPRecipes(ctx *middleware.Context, handler ReloadDSPRecipesHandler) *ReloadDSPRecipes {
	return &ReloadDSPRecipes{Context: ctx, Handler: handler}
}

/* ReloadDSPRecipes swagger:route POST /dsp/recipes/reload reloadDSPRecipes

Re-scrapes the DSP recipes

Re-scrapes the DSP recipes

*/
type ReloadDSPRecipes struct {
	Context *middleware.Context
	Handler ReloadDSPRecipesHandler
}

func (o *ReloadDSPRecipes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewReloadDSPRecipesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
