package lib

import (
	"encoding/json"
	"net/http"
	"project/model"
)

func JsonResponse(w http.ResponseWriter) Response {
	w.Header().Set("Content-Type", "application/json")
	return Response{W: w}
}

type Response struct {
	W http.ResponseWriter
}

func (response Response) Fail(statusCode int, message string) {
	r := model.Response{
		Success:    false,
		StatusCode: &statusCode,
		Message:    &message,
	}

	if statusCode == 0 {
		r.StatusCode = nil
	}

	json.NewEncoder(response.W).Encode(r)
}

func (response Response) Success(statusCode int, message string, data interface{}) {
	r := model.Response{
		Success:    true,
		StatusCode: &statusCode,
		Message:    &message,
		Data:       data,
	}

	if statusCode == 0 {
		r.StatusCode = nil
	}

	if message == "" {
		r.Message = nil
	}

	json.NewEncoder(response.W).Encode(r)
}
