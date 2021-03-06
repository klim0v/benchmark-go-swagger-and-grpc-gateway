// Code generated by go-swagger; DO NOT EDIT.

package shop

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ShopListProductsHandlerFunc turns a function with the right signature into a shop list products handler
type ShopListProductsHandlerFunc func(ShopListProductsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShopListProductsHandlerFunc) Handle(params ShopListProductsParams) middleware.Responder {
	return fn(params)
}

// ShopListProductsHandler interface for that can handle valid shop list products params
type ShopListProductsHandler interface {
	Handle(ShopListProductsParams) middleware.Responder
}

// NewShopListProducts creates a new http.Handler for the shop list products operation
func NewShopListProducts(ctx *middleware.Context, handler ShopListProductsHandler) *ShopListProducts {
	return &ShopListProducts{Context: ctx, Handler: handler}
}

/*ShopListProducts swagger:route GET /products Shop shopListProducts

ShopListProducts shop list products API

*/
type ShopListProducts struct {
	Context *middleware.Context
	Handler ShopListProductsHandler
}

func (o *ShopListProducts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShopListProductsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
