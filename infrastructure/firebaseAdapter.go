package infrastructure

import (
	"context"
	"fmt"
	"personal-page-back/domain"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	ProjectIdKey     = "FIREBASE_PROJECT_ID"
	skillsCollection = "/skills/"
)

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

	f.client.NewRef(skillsCollection).Get(ctx, &m)

	return mapToSkillArray(m)
}

func (f *firebaseAdapter) GetSkillByName(ctx context.Context, name string) *domain.Skill {

	m := make(map[string]interface{})

	f.client.NewRef(skillsCollection+name).Get(ctx, &m)

	if m == nil {
		return nil
	} else {
		skill := mapSingleSkill(m)
		return &skill
	}

}

func (f *firebaseAdapter) UpdateSkill(ctx context.Context, s *domain.Skill) error {

	err := f.client.NewRef(skillsCollection+s.Name).Set(ctx, s)
	return err
}

func (f *firebaseAdapter) IncrementSkill(ctx context.Context, s domain.Skill) {

}

func mapToSkillArray(values firebaseSkills) []domain.Skill {
	array := make([]domain.Skill, 0)

	for _, v := range values {

		skill := mapSingleSkill(v)
		array = append(array, skill)
	}

	return array
}

func mapSingleSkill(value map[string]interface{}) domain.Skill {

	name := value["name"]
	count := value["count"]
	return domain.Skill{Name: name.(string), Count: count.(float64)}
}
