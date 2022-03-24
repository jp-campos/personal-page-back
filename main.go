package main

import (
	"net/http"
	"personal-page-back/domain"
	"personal-page-back/entrypoints"
	"personal-page-back/infrastructure"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/skills", entrypoints.GetSkill).Queries("prefix", "{prefix}").Methods(http.MethodGet)
	r.HandleFunc("/skills", entrypoints.GetSkills).Methods(http.MethodGet)
	r.HandleFunc("/skill", entrypoints.PostSkill).Methods(http.MethodPost)
	return r
}

func main() {

	domain.InitRepository(infrastructure.NewFirebaseAdapter("url"))

	r := initRouter()
	http.ListenAndServe(":8080", r)

}
