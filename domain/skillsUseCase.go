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

	for _, e := range allSkills {
		if strings.HasPrefix(e.Name, strings.ToTitle(prefix)) {
			filteredSkills = append(filteredSkills, e)
		}
	}

	return filteredSkills
}
