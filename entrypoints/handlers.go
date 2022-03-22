package entrypoints

import (
	"encoding/json"
	"net/http"
	"personal-page-back/domain"

	"github.com/gorilla/mux"
)

func GetSkills(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	skills := domain.GetSkills()

	json.NewEncoder(w).Encode(skills)
}

func GetSkill(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	value, _ := vars["id"]

	skills := domain.GetSkillsStartingWith(value)

	json.NewEncoder(w).Encode(skills)
}

func PostSkill(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("posteo"))
}
