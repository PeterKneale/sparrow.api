//************************************************************************//
// API "Gateway API": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/mango/gateway/design
// --out=C:\dev\go\src\github.com\simplicate\mango\gateway\temp
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
	service.Mux.Handle("GET", "/account", ctrl.MuxHandler("Read", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Read", "route", "GET /account")

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
	service.Mux.Handle("PUT", "/account", ctrl.MuxHandler("Update", h, unmarshalUpdateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Update", "route", "PUT /account")
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
	service.Mux.Handle("POST", "/users", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /users")

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
	service.Mux.Handle("DELETE", "/users/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /users/:id")

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
	service.Mux.Handle("GET", "/users", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /users")

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
	service.Mux.Handle("GET", "/users/:id", ctrl.MuxHandler("Read", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Read", "route", "GET /users/:id")

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
	service.Mux.Handle("PUT", "/users/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "PUT /users/:id")
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
	service.Mux.Handle("GET", "/alive", ctrl.MuxHandler("Alive", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Alive", "route", "GET /alive")
}

// Js1Controller is the controller interface for the Js1 actions.
type Js1Controller interface {
	goa.Muxer
	goa.FileServer
}

// MountJs1Controller "mounts" a Js1 resource controller on the given service.
func MountJs1Controller(service *goa.Service, ctrl Js1Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/js1/*filepath", ctrl.MuxHandler("preflight", handleJs1Origin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/js1/*filepath", "web1/js1")
	h = handleJs1Origin(h)
	service.Mux.Handle("GET", "/js1/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js1", "files", "web1/js1", "route", "GET /js1/*filepath")

	h = ctrl.FileHandler("/js1/", "web1\\js1\\index.html")
	h = handleJs1Origin(h)
	service.Mux.Handle("GET", "/js1/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js1", "files", "web1\\js1\\index.html", "route", "GET /js1/")
}

// handleJs1Origin applies the CORS response headers corresponding to the origin.
func handleJs1Origin(h goa.Handler) goa.Handler {

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

// Js2Controller is the controller interface for the Js2 actions.
type Js2Controller interface {
	goa.Muxer
	goa.FileServer
}

// MountJs2Controller "mounts" a Js2 resource controller on the given service.
func MountJs2Controller(service *goa.Service, ctrl Js2Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/js2/*filepath", ctrl.MuxHandler("preflight", handleJs2Origin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/js2/*filepath", "web2/js2")
	h = handleJs2Origin(h)
	service.Mux.Handle("GET", "/js2/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js2", "files", "web2/js2", "route", "GET /js2/*filepath")

	h = ctrl.FileHandler("/js2/", "web2\\js2\\index.html")
	h = handleJs2Origin(h)
	service.Mux.Handle("GET", "/js2/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js2", "files", "web2\\js2\\index.html", "route", "GET /js2/")
}

// handleJs2Origin applies the CORS response headers corresponding to the origin.
func handleJs2Origin(h goa.Handler) goa.Handler {

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

	h = ctrl.FileHandler("/swagger.json", "web1/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "web1/swagger/swagger.json", "route", "GET /swagger.json")
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

// Web1Controller is the controller interface for the Web1 actions.
type Web1Controller interface {
	goa.Muxer
	goa.FileServer
}

// MountWeb1Controller "mounts" a Web1 resource controller on the given service.
func MountWeb1Controller(service *goa.Service, ctrl Web1Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/web1", ctrl.MuxHandler("preflight", handleWeb1Origin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/web1", "web1/index.html")
	h = handleWeb1Origin(h)
	service.Mux.Handle("GET", "/web1", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Web1", "files", "web1/index.html", "route", "GET /web1")
}

// handleWeb1Origin applies the CORS response headers corresponding to the origin.
func handleWeb1Origin(h goa.Handler) goa.Handler {

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

// Web2Controller is the controller interface for the Web2 actions.
type Web2Controller interface {
	goa.Muxer
	goa.FileServer
}

// MountWeb2Controller "mounts" a Web2 resource controller on the given service.
func MountWeb2Controller(service *goa.Service, ctrl Web2Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/web2", ctrl.MuxHandler("preflight", handleWeb2Origin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/web2", "web2/index.html")
	h = handleWeb2Origin(h)
	service.Mux.Handle("GET", "/web2", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Web2", "files", "web2/index.html", "route", "GET /web2")
}

// handleWeb2Origin applies the CORS response headers corresponding to the origin.
func handleWeb2Origin(h goa.Handler) goa.Handler {

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
