package apiserver

import "net/http"

func (as *ApiServer) routes() http.Handler {
	router := http.NewServeMux()

	return router
}
