package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error,omitempty"`
	} `json:"status"`
}

func NewResponse(success bool, message string, err error) Response {
	resp := Response{}
	resp.Status.Success = success
	resp.Status.Message = message
	if err != nil {
		resp.Status.Error = err.Error()
	}
	return resp
}

func JsonResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	_, _ = w.Write(response)
}
