package httpapi

import "github.com/msamad/group-events/backend/internal/sdui"

type Response struct {
	Data interface{}       `json:"data"`
	UI   sdui.UIDescriptor `json:"ui"`
}

type ErrorResponse struct {
	Error APIError          `json:"error"`
	UI    sdui.UIDescriptor `json:"ui,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}