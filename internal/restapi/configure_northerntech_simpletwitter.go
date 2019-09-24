// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
)

//go:generate swagger generate server --target ../../internal --name NortherntechSimpletwitter --spec ../../api/swagger/swagger.yml --exclude-main

func configureFlags(api *operations.NortherntechSimpletwitterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NortherntechSimpletwitterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DeleteTweetsIDHandler == nil {
		api.DeleteTweetsIDHandler = operations.DeleteTweetsIDHandlerFunc(func(params operations.DeleteTweetsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeleteTweetsID has not yet been implemented")
		})
	}
	if api.GetHealthHandler == nil {
		api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetHealth has not yet been implemented")
		})
	}
	if api.GetTweetsHandler == nil {
		api.GetTweetsHandler = operations.GetTweetsHandlerFunc(func(params operations.GetTweetsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetTweets has not yet been implemented")
		})
	}
	if api.GetTweetsIDHandler == nil {
		api.GetTweetsIDHandler = operations.GetTweetsIDHandlerFunc(func(params operations.GetTweetsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetTweetsID has not yet been implemented")
		})
	}
	if api.PatchTweetsHandler == nil {
		api.PatchTweetsHandler = operations.PatchTweetsHandlerFunc(func(params operations.PatchTweetsParams) middleware.Responder {
			return middleware.NotImplemented("operation .PatchTweets has not yet been implemented")
		})
	}
	if api.PostTweetsHandler == nil {
		api.PostTweetsHandler = operations.PostTweetsHandlerFunc(func(params operations.PostTweetsParams) middleware.Responder {
			return middleware.NotImplemented("operation .PostTweets has not yet been implemented")
		})
	}

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
