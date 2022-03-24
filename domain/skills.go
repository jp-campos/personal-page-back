package domain

import "context"

type Skill struct {
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

type SkillGateWay interface {
	GetSkills(context.Context) []Skill
}
