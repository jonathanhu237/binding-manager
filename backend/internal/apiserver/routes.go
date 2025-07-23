package apiserver

import "net/http"

func (as *ApiServer) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/health-check", as.healthCheckHandler)

	return router
}
