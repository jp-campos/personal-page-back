package entrypoints

import (
	"encoding/json"
	"net/http"
	"personal-page-back/domain"
)

const (
	contentTypeHeader = "Content-Type"
	corsHeader        = "Access-Control-Allow-Origin"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set(contentTypeHeader, "application/json")
	w.Header().Set(corsHeader, "http://localhost:3000")

}

func GetSkills(w http.ResponseWriter, req *http.Request) {
	setHeaders(w)

	skills := domain.GetSkills(req.Context())

	json.NewEncoder(w).Encode(skills)
}

func GetSkill(w http.ResponseWriter, req *http.Request) {

	setHeaders(w)
	prefix := req.FormValue("prefix")

	skills := domain.GetSkillsStartingWith(req.Context(), prefix)

	json.NewEncoder(w).Encode(skills)
}

func PostSkill(w http.ResponseWriter, req *http.Request) {

	m := make(map[string]string)
	json.NewDecoder(req.Body).Decode(&m)

	domain.IncrementSkill(req.Context(), m["skill"])

}
