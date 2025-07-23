package apiserver

import "net/http"

func (as *ApiServer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": as.config.Server.Environment,
		},
	}

	as.successResponse(w, r, &data)
}
