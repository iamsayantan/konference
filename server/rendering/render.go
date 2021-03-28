package rendering

import (
	"encoding/json"
	chimw "github.com/go-chi/chi/middleware"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

func RenderSuccess(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
	}

	renderJSON(w, r, statusCode, resp)
}

func RenderSuccessWithData(w http.ResponseWriter, r *http.Request, message string, statusCode int, data interface{}) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	renderJSON(w, r, statusCode, resp)
}

func RenderError(w http.ResponseWriter, r *http.Request, message, errorCode string, statusCode int) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
		ErrorCode:  errorCode,
	}

	renderJSON(w, r, statusCode, resp)
}

func RenderErrorsWithData(w http.ResponseWriter, r *http.Request, message, errorCode string, statusCode int, errorData interface{}) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
		ErrorCode:  errorCode,
		Errors:     errorData,
	}

	renderJSON(w, r, statusCode, resp)
}

func renderJSON(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(chimw.RequestIDHeader, chimw.GetReqID(r.Context()))

	w.WriteHeader(statusCode)
	resp, _ := json.Marshal(v)
	_, _ = w.Write(resp)
}
