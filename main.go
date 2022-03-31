package main

import (
	"context"
	"log"
	"net/http"
	"personal-page-back/domain"
	"personal-page-back/entrypoints"
	"personal-page-back/infrastructure"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(entrypoints.Options).Methods(http.MethodOptions)
	r.HandleFunc("/skills", entrypoints.GetSkillsWithPrefix).Queries("prefix", "{prefix}").Methods(http.MethodGet)
	r.HandleFunc("/skills", entrypoints.GetSkills).Methods(http.MethodGet)
	r.HandleFunc("/skill", entrypoints.PostSkill).Methods(http.MethodPost)
	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	domain.InitSkillRepository(infrastructure.NewFirebaseAdapter(ctx))

	r := initRouter()
	http.ListenAndServe(":8080", r)

}
