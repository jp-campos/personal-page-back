package entrypoints

import (
	"encoding/json"
	"net/http"
	"personal-page-back/domain"
)

const (
	contentTypeHeader  = "Content-Type"
	corsOrigin         = "Access-Control-Allow-Origin"
	allowedCorsMethods = "Access-Control-Allow-Methods"
	allowedCorsHeaders = "Access-Control-Allow-Headers"
	corsMaxAge         = "Access-Control-Max-Age"
	authHeader         = "Authorization"
	contectType        = "Content-Type"
	acceptHeader       = "Accept"
	origin             = "Origin"
	userAgent          = "User-Agent"
)

var allowedHeaders = [...]string{"Authorization", "Content-Type", "Accept", "Origin", "User-Agent"}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set(contentTypeHeader, "application/json")
	w.Header().Set(corsOrigin, "*")

}

func GetSkills(w http.ResponseWriter, req *http.Request) {
	setHeaders(w)

	skills := domain.GetSkills(req.Context())

	json.NewEncoder(w).Encode(skills)
}

func GetSkillsWithPrefix(w http.ResponseWriter, req *http.Request) {

	setHeaders(w)
	prefix := req.FormValue("prefix")

	skills := domain.GetSkillsStartingWith(req.Context(), prefix)

	json.NewEncoder(w).Encode(skills)
}

func PostSkill(w http.ResponseWriter, req *http.Request) {
	setHeaders(w)
	m := make(map[string]string)
	json.NewDecoder(req.Body).Decode(&m)

	domain.IncrementSkill(req.Context(), m["skill"])
}

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
