package domain

import (
	"context"
	"fmt"
	"strings"
)

var repository SkillGateWay

func InitRepository(repo SkillGateWay) {
	repository = repo

}

func GetSkills(ctx context.Context) []Skill {

	return repository.GetSkills(ctx)
}

func GetSkillsStartingWith(ctx context.Context, prefix string) []Skill {

	allSkills := repository.GetSkills(ctx)
	fmt.Printf("Size in usecase %v \n", len(allSkills))
	filteredSkills := make([]Skill, 0)

	for _, e := range allSkills {
		if len(filteredSkills) == 5 {
			break
		}
		normalizedSavedSkill := strings.ToUpper(e.Name)
		normalizedQuery := strings.ToUpper(prefix)
		if strings.Contains(normalizedSavedSkill, normalizedQuery) {
			filteredSkills = append(filteredSkills, e)

		}
	}

	return filteredSkills
}
