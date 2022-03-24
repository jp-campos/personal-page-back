package domain

import (
	"strings"
)

var repository SkillGateWay

func InitRepository(repo SkillGateWay) {
	repository = repo
}

func GetSkills() []Skill {

	return repository.GetSkills()
}

func GetSkillsStartingWith(prefix string) []Skill {

	allSkills := repository.GetSkills()

	filteredSkills := make([]Skill, 0)
	for i, e := range allSkills {
		if i == 5 {
			break
		}
		normalizedSavedSkill := strings.ToUpper(e.Name)
		normalizedQuery := strings.ToUpper(prefix)
		if strings.HasPrefix(normalizedSavedSkill, normalizedQuery) {
			filteredSkills = append(filteredSkills, e)
		}
	}

	return filteredSkills
}
