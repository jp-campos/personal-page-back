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

	s, err := domain.IncrementSkill(req.Context(), m["skill"])

	if err != nil {
		w.WriteHeader(500)
	} else {
		json.NewEncoder(w).Encode(s)
	}
}
