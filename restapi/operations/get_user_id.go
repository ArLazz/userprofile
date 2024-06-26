// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserIDHandlerFunc turns a function with the right signature into a get user ID handler
type GetUserIDHandlerFunc func(GetUserIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserIDHandlerFunc) Handle(params GetUserIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUserIDHandler interface for that can handle valid get user ID params
type GetUserIDHandler interface {
	Handle(GetUserIDParams, interface{}) middleware.Responder
}

// NewGetUserID creates a new http.Handler for the get user ID operation
func NewGetUserID(ctx *middleware.Context, handler GetUserIDHandler) *GetUserID {
	return &GetUserID{Context: ctx, Handler: handler}
}

/*
	GetUserID swagger:route GET /user/{id} getUserId

Get user by ID
*/
type GetUserID struct {
	Context *middleware.Context
	Handler GetUserIDHandler
}

func (o *GetUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
