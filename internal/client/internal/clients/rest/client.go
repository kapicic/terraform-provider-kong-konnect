package rest

import (
	"github.com/kong-sdk/internal/clients/rest/handlers"
	"github.com/kong-sdk/internal/clients/rest/hooks"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
)

type RestClient struct {
	handlers *handlers.HandlerChain
}

func NewRestClient(baseUrl string, bearerToken string) *RestClient {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler()
	bearerTokenHandler := handlers.NewBearerTokenHandler(bearerToken)
	hookHandler := handlers.NewHookHandler(hooks.NewDefaultHook())
	terminatingHandler := handlers.NewTerminatingHandler()

	handlers := handlers.BuildHandlerChain().
		AddHandler(defaultHeadersHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient{
		handlers: handlers,
	}
}

func (client *RestClient) Call(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	return client.handlers.CallApi(request)
}
