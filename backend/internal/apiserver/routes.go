package apiserver

import "net/http"

func (as *ApiServer) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/server-info", as.getServerInfo)

	return router
}
