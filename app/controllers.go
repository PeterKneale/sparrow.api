//************************************************************************//
// API "Sparrow API": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/sparrow.api/design
// --out=$(GOPATH)\src\github.com\Simplicate\sparrow.api\
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Muxer
	Read(*ReadAccountContext) error
	Update(*UpdateAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service *goa.Service, ctrl AccountController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/account", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReadAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Read(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("GET", "/api/account", ctrl.MuxHandler("Read", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Read", "route", "GET /api/account")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("PUT", "/api/account", ctrl.MuxHandler("Update", h, unmarshalUpdateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Update", "route", "PUT /api/account")
}

// handleAccountOrigin applies the CORS response headers corresponding to the origin.
func handleAccountOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateAccountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Read(*ReadUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/users", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/users/:id", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/api/users", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /api/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("DELETE", "/api/users/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/users/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/api/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /api/users")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReadUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Read(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/api/users/:id", ctrl.MuxHandler("Read", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Read", "route", "GET /api/users/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("PUT", "/api/users/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PUT /api/users/:id")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Alive(*AliveHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/alive", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAliveHealthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Alive(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/api/alive", ctrl.MuxHandler("Alive", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Alive", "route", "GET /api/alive")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.yaml", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger.yaml", "swagger/swagger.yaml")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.yaml", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.yaml", "route", "GET /swagger.yaml")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
