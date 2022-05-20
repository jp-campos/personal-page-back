package infrastructure

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"personal-page-back/domain"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	ProjectIdKey                 = "FIREBASE_PROJECT_ID"
	skillsCollection             = "/skills/"
	unclassifiedSkillsCollection = "/unclassified-skills/"
)

type firebaseAdapter struct {
	client *db.Client
}

type firebaseSkills map[string]map[string]interface{}

func NewFirebaseAdapter(ctx context.Context) *firebaseAdapter {
	fileName := "serviceAccountKey.json"
	replaceSecretsServiceAccout(fileName)
	opt := option.WithCredentialsFile(fileName)

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
	fmt.Println("Created firebase client", client)
	return &firebaseAdapter{client: client}
}

func (f *firebaseAdapter) Skills(ctx context.Context) []domain.Skill {

	m := make(firebaseSkills)

	f.client.NewRef(skillsCollection).Get(ctx, &m)

	return mapToSkillArray(m)
}

func (f *firebaseAdapter) SkillByName(ctx context.Context, name string) *domain.Skill {
	return f.getSkillInCollection(ctx, name, skillsCollection)
}

func (f *firebaseAdapter) UnclassifiedSkillByName(ctx context.Context, name string) *domain.Skill {
	skillRoute := unclassifiedEncodedName(name)
	return f.getSkillInCollection(ctx, skillRoute, unclassifiedSkillsCollection)
}

func (f *firebaseAdapter) UpdateSkill(ctx context.Context, s *domain.Skill) error {
	return f.client.NewRef(skillsCollection+s.Name).Set(ctx, s)
}

func (f *firebaseAdapter) UpdateUnclassifiedSkill(ctx context.Context, s *domain.Skill) error {

	skillRoute := unclassifiedEncodedName(s.Name)
	return f.client.NewRef(unclassifiedSkillsCollection+skillRoute).Set(ctx, s)

}

func (f *firebaseAdapter) CreateUnclassifiedSkill(ctx context.Context, s domain.Skill) error {
	skillRoute := unclassifiedEncodedName(s.Name)

	return f.client.NewRef(unclassifiedSkillsCollection+skillRoute).Set(ctx, s)
}

func (f *firebaseAdapter) getSkillInCollection(ctx context.Context, name, collection string) *domain.Skill {

	m := make(map[string]interface{})
	f.client.NewRef(collection+name).Get(ctx, &m)

	if m == nil || len(m) == 0 {
		return nil
	} else {
		skill := mapSingleSkill(m)
		return &skill
	}
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

func replaceSecretsServiceAccout(fileName string) {

	read, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	privateKeyReplace := "${privateKey}"
	privateKeyEnvKey := "PRIVATE_KEY"
	newFileText := strings.Replace(string(read), privateKeyReplace, os.Getenv(privateKeyEnvKey), 1)

	err = ioutil.WriteFile(fileName, []byte(newFileText), 0)

	if err != nil {
		fmt.Println(err)
	}

}
func unclassifiedEncodedName(name string) string {
	return base64.StdEncoding.EncodeToString([]byte(name))
}
