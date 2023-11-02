package handlers

import (
	"errors"
	"fmt"

	"github.com/kong-sdk/internal/clients/rest/httptransport"
)

type BearerTokenHandler struct {
	bearerToken string
	nextHandler Handler
}

func NewBearerTokenHandler(bearerToken string) *BearerTokenHandler {
	return &BearerTokenHandler{
		bearerToken: bearerToken,
		nextHandler: nil,
	}
}

func (h *BearerTokenHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	nextRequest := request.Clone()
	nextRequest.AddHeader("Authorization", fmt.Sprintf("Bearer %s", h.bearerToken))

	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.nextHandler.Handle(nextRequest)
}

func (h *BearerTokenHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
