package domain

import "context"

type Skill struct {
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

type SkillGateWay interface {
	Skills(context.Context) []Skill
	UpdateSkill(context.Context, *Skill) error
	SkillByName(context.Context, string) *Skill
	UnclassifiedSkillByName(context.Context, string) *Skill
	CreateUnclassifiedSkill(context.Context, Skill) error
	UpdateUnclassifiedSkill(context.Context, *Skill) error
}
