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

	//vars := mux.Vars(req)
	//value, _ := vars["id"]
	prefix := req.FormValue("prefix")
	println(prefix)

	skills := domain.GetSkillsStartingWith(req.Context(), prefix)

	json.NewEncoder(w).Encode(skills)
}

func PostSkill(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("posteo"))
}
