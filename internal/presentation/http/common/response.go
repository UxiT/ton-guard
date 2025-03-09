package common

import (
	"decard/internal/domain"
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  any   `json:"data"`
	Items []any `json:"items"`
}

type ErrorResponse struct {
	Code   domain.ErrorCode `json:"code"`
	Errors any              `json:"errors"`
}

func JSONResponse(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(Response{Data: v})
}

func JSONErrorResponse(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")

	var response ErrorResponse

	if apiErr, ok := err.(domain.ApplicationError); ok {
		response = ErrorResponse{
			Code: apiErr.Code,
		}

		w.WriteHeader(apiErr.HTTPCode)
	} else {
		response = ErrorResponse{
			Code: domain.InternalError,
		}

		w.WriteHeader(http.StatusInternalServerError)
	}

	return json.NewEncoder(w).Encode(response)
}
