// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/klim0v/benchmark-go-swagger-and-grpc-gateway/go-swagger/models"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/klim0v/benchmark-go-swagger-and-grpc-gateway/go-swagger/restapi/operations"
	"github.com/klim0v/benchmark-go-swagger-and-grpc-gateway/go-swagger/restapi/operations/shop"
)

//go:generate swagger generate server --target ../../go-swagger --name ShopShopProto --spec ../../grpc-gateway/api.swagger.json --principal interface{}

func configureFlags(api *operations.ShopShopProtoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ShopShopProtoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ShopShopGetProductHandler = shop.ShopGetProductHandlerFunc(func(params shop.ShopGetProductParams) middleware.Responder {
		return shop.NewShopGetProductOK().WithPayload(&models.ShopProductResponse{
			Title:    "A",
			Favorite: true,
			Price:    1999.99,
			Count:    "200",
			Image:    "https://image",
			Vote: []*models.ProductResponseVote{
				{
					Name: "A",
					Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
				},
			},
		})
	})
	api.ShopShopListProductsHandler = shop.ShopListProductsHandlerFunc(func(params shop.ShopListProductsParams) middleware.Responder {
		return shop.NewShopListProductsOK().WithPayload(&models.ShopListProductsResponse{Products: []*models.ShopProductResponse{
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
			{
				Title:    "A",
				Favorite: true,
				Price:    1999.99,
				Count:    "200",
				Image:    "https://image",
				Vote: []*models.ProductResponseVote{
					{
						Name: "A",
						Text: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit...",
					},
				},
			},
		}})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
