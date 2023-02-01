// Code generated by go-swagger; DO NOT EDIT.

package pets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PetsListHandlerFunc turns a function with the right signature into a pets list handler
type PetsListHandlerFunc func(PetsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PetsListHandlerFunc) Handle(params PetsListParams) middleware.Responder {
	return fn(params)
}

// PetsListHandler interface for that can handle valid pets list params
type PetsListHandler interface {
	Handle(PetsListParams) middleware.Responder
}

// NewPetsList creates a new http.Handler for the pets list operation
func NewPetsList(ctx *middleware.Context, handler PetsListHandler) *PetsList {
	return &PetsList{Context: ctx, Handler: handler}
}

/*
	PetsList swagger:route GET /pets pets petsList

# Get list of currently available pets

It returns a list of nested objects which contains all pets and their qualities
*/
type PetsList struct {
	Context *middleware.Context
	Handler PetsListHandler
}

func (o *PetsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPetsListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
