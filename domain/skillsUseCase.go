package domain

import (
	"context"
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

func IncrementSkill(ctx context.Context, name string) error {

	skill := repository.GetSkillByName(ctx, name)

	var err error
	if skill == nil {
		incrementUnclassifiedSkill(ctx, name)
	} else {
		skill.Count++
		err = repository.UpdateSkill(ctx, skill)

	}
	return err
}

func incrementUnclassifiedSkill(ctx context.Context, name string) {

	upperCaseName := strings.ToUpper(name)
	skill := repository.GetUnclassifiedSkillByName(ctx, upperCaseName)

	if skill == nil {
		newSkill := Skill{Name: upperCaseName, Count: 1}
		repository.CreateUnclassifiedSkill(ctx, newSkill)

	} else {
		skill.Count++
		repository.UpdateUnclassifiedSkill(ctx, skill)
	}

}
