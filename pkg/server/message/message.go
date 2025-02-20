// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Message defines model for Message.
type Message struct {
	Author    *string `json:"author,omitempty"`
	Content   *string `json:"content,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// MessageCreate defines model for MessageCreate.
type MessageCreate struct {
	Content string `json:"content"`
}

// SendMessageToChatRoomJSONRequestBody defines body for SendMessageToChatRoom for application/json ContentType.
type SendMessageToChatRoomJSONRequestBody = MessageCreate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get chat room messages
	// (GET /rooms/{id}/messages)
	GetChatRoomMessages(ctx echo.Context, id string) error
	// Send a message to a chat room
	// (POST /rooms/{id}/messages)
	SendMessageToChatRoom(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetChatRoomMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetChatRoomMessages(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{"email:w"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetChatRoomMessages(ctx, id)
	return err
}

// SendMessageToChatRoom converts echo context to params.
func (w *ServerInterfaceWrapper) SendMessageToChatRoom(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{"email:w"})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SendMessageToChatRoom(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/rooms/:id/messages", wrapper.GetChatRoomMessages)
	router.POST(baseURL+"/rooms/:id/messages", wrapper.SendMessageToChatRoom)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7STT2/UMBDFv4o1cEybBW6+tZVARVRC7UocVntwk9mNq9jj2hNQFOW7o3GyfwqpBEjc",
	"vJ7ZeS+/eR6gIhfIo+cEeoBUNehMPt5hSmaPcgyRAka2mAum44ainLgPCBoSR+v3MBZQkWf0vFhj6zCx",
	"cWGhKuXphh6fsGLpn/VvIhpecPG61FhAxOfORqxBb46N2980xgISVl203IPeDHCNJmK86rgBvdmO21P5",
	"QbBMuudNAzzmXx8pOsOg4fO3NRQTRBGaqnAUbpgDjKJr/Y6ydcutVO7RtBdr61DdNIbV1ddbdaFmAuqO",
	"6q5FKOA7xmTJg4Z3l6vLlVCigN4ECxo+5KsCguEmOy0jkUvlYOuxdNOofL/HTE1gGrbkb2vQ8AlZlO+J",
	"3N2hV2ZF45AxpgzIirTMhwK8yZ9oazjHzbHDGYBZWs1WmlMgnyYv71erX5ZpQmhtlY2VT0k+djibZxld",
	"/uPbiDvQ8KY85becw1sekntKlYnR9BP5GlMVbeCJ4xebWNFOHfm8EgoJEjpjW/0DpmR0zpnYT+RUJUsT",
	"3GeDCgiUFkg/oK9nh2s6MP9vrJ87THxNdf9XmP+A7vwux5fPTTyNyzt+Cf6Q7YSeVeqqClPadW3b/8MG",
	"hKgyB/SKSZnTRsTh+DMAAP//L52wTOcEAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
