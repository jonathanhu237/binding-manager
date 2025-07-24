package apiserver

import "net/http"

func (as *ApiServer) getServerInfo(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": as.cfg.ApiServer.Environment,
		},
	}

	as.successResponse(w, r, &data)
}
