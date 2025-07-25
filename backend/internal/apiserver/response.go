package apiserver

import (
	"net/http"

	"github.com/jonathanhu237/binding-manager/backend/internal/unierror"
)

type envelope map[string]any

type unifiedResponse struct {
	Success bool                   `json:"success"`
	Data    *envelope              `json:"data"`
	Error   *unierror.UnifiedError `json:"error"`
}

func (as *ApiServer) successResponse(w http.ResponseWriter, r *http.Request, data *envelope) {
	resp := unifiedResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}

	if err := as.writeJson(w, http.StatusOK, resp, nil); err != nil {
		as.internalServerError(w, r, err)
	}
}

func (as *ApiServer) logError(r *http.Request, err error) {
	as.logger.Error(err.Error(), "method", r.Method, "url", r.URL.String())
}

func (as *ApiServer) errorResponse(w http.ResponseWriter, r *http.Request, err *unierror.UnifiedError) {
	resp := unifiedResponse{
		Success: false,
		Data:    nil,
	}

	var unifiedErr *unierror.UnifiedError
	if errors.As(err, &unifiedErr) {
		resp.Error = unifiedErr
	} else {
		resp.Error = &unierror.UnifiedError{
			Code:    status,
			Message: err.Error(),
			Details: nil,
		}
	}

	if err := as.writeJson(w, err.Code, resp, nil); err != nil {
		as.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (as *ApiServer) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	as.logError(r, err)
	as.errorResponse(w, r, unierror.ErrInternalServerError)
}
