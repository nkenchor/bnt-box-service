package api

import (
	ports "bnt/bnt-box-service/internal/port"
)

type HTTPHandler struct {
	boxService ports.BoxService
}

func NewHTTPHandler(
	prefixService ports.BoxService) *HTTPHandler {
	return &HTTPHandler{
		boxService: prefixService,
	}
}
