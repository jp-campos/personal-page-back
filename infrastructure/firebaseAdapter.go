package infrastructure

import (
	"context"
	"fmt"
	"personal-page-back/domain"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const ProjectIdKey = "FIREBASE_PROJECT_ID"

type firebaseAdapter struct {
	client *db.Client
}

type firebaseSkills map[string]map[string]interface{}

func NewFirebaseAdapter(ctx context.Context) *firebaseAdapter {

	opt := option.WithCredentialsFile("serviceAccountKey.json")

	conf := &firebase.Config{
		DatabaseURL: "https://personal-page-6d1ac-default-rtdb.firebaseio.com/"}
	app, err := firebase.NewApp(ctx, conf, opt)

	if err != nil {
		err = fmt.Errorf("Error creating app: %v", err)
		fmt.Println(err)
	}

	client, err := app.Database(ctx)

	if err != nil {
		err = fmt.Errorf("Error creating client: %v", err)
		fmt.Println(err)
	}

	return &firebaseAdapter{client: client}
}

func (f *firebaseAdapter) GetSkills(ctx context.Context) []domain.Skill {

	m := make(firebaseSkills)

	f.client.NewRef("/skills").Get(ctx, &m)

	return mapToSkillArray(m)
}

func (f *firebaseAdapter) IncrementNewSkill(s domain.Skill) {

}

func (f *firebaseAdapter) IncrementSkill(s domain.Skill) {

}

func mapToSkillArray(values firebaseSkills) []domain.Skill {
	array := make([]domain.Skill, 0)

	for _, v := range values {

		name := v["name"]
		count := v["count"]

		array = append(array, domain.Skill{Name: name.(string), Count: count.(float64)})
	}

	return array
}
