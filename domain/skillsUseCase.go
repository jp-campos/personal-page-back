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

	return skillRepo.Skills(ctx)
}

func GetSkillsStartingWith(ctx context.Context, prefix string) []Skill {

	allSkills := skillRepo.Skills(ctx)
	filteredSkills := make([]Skill, 0, 5)

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

func IncrementSkill(ctx context.Context, name string) (Skill, error) {

	skill, err := skillRepo.SkillByName(ctx, name)

	if err != nil {
		return Skill{}, err
	}
	if skill == nil {
		return incrementUnclassifiedSkill(ctx, name)
	} else {
		skill.Count++
		err = skillRepo.UpdateSkill(ctx, skill)
		return *skill, err
	}

}

func incrementUnclassifiedSkill(ctx context.Context, name string) (Skill, error) {

	upperCaseName := strings.ToUpper(name)
	skill, err := skillRepo.UnclassifiedSkillByName(ctx, upperCaseName)

	if err != nil {
		return Skill{}, err
	}

	if skill == nil {
		newSkill := Skill{Name: upperCaseName, Count: 1}
		err = skillRepo.CreateUnclassifiedSkill(ctx, newSkill)
		skill = &newSkill

	} else {
		skill.Count++
		err = skillRepo.UpdateUnclassifiedSkill(ctx, skill)

	}

	return *skill, err
}
