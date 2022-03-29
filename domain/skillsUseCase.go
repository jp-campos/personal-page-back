package domain

import (
	"context"
	"strings"
)

var skillRepo SkillGateWay

func InitSkillRepository(repo SkillGateWay) {
	skillRepo = repo
}

func GetSkills(ctx context.Context) []Skill {

	return skillRepo.GetSkills(ctx)
}

func GetSkillsStartingWith(ctx context.Context, prefix string) []Skill {

	allSkills := skillRepo.GetSkills(ctx)
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

	skill := skillRepo.GetSkillByName(ctx, name)

	var err error
	if skill == nil {
		incrementUnclassifiedSkill(ctx, name)
	} else {
		skill.Count++
		err = skillRepo.UpdateSkill(ctx, skill)

	}
	return err
}

func incrementUnclassifiedSkill(ctx context.Context, name string) {

	upperCaseName := strings.ToUpper(name)
	skill := skillRepo.GetUnclassifiedSkillByName(ctx, upperCaseName)

	if skill == nil {
		newSkill := Skill{Name: upperCaseName, Count: 1}
		skillRepo.CreateUnclassifiedSkill(ctx, newSkill)

	} else {
		skill.Count++
		skillRepo.UpdateUnclassifiedSkill(ctx, skill)
	}

}
