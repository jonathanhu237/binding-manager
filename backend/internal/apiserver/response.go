package apiserver

import "net/http"

type envelope map[string]any

type unifiedResponse struct {
	Success bool          `json:"success"`
	Data    *envelope     `json:"data"`
	Error   *ServiceError `json:"error"`
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

func (as *ApiServer) errorResponse(w http.ResponseWriter, r *http.Request, serviceError ServiceError) {
	resp := unifiedResponse{
		Success: false,
		Data:    nil,
		Error:   &serviceError,
	}

	if err := as.writeJson(w, serviceError.Code, resp, nil); err != nil {
		as.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (as *ApiServer) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	as.logError(r, err)
	as.errorResponse(w, r, ErrInternalServerError)
}
