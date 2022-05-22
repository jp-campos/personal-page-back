package infrastructure

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"personal-page-back/domain"
	"strings"
)

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