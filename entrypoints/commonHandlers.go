package entrypoints

import "net/http"

func Options(w http.ResponseWriter, req *http.Request) {
	w.Header().Add(allowedCorsMethods, http.MethodGet)
	w.Header().Add(allowedCorsMethods, http.MethodPost)
	w.Header().Add(allowedCorsMethods, http.MethodOptions)
	w.Header().Add(corsMaxAge, "500")
	for _, e := range allowedHeaders {
		w.Header().Add(allowedCorsHeaders, e)
	}
	w.Header().Set(corsOrigin, "*")
}
