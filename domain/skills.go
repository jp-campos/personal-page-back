package domain

type Skill struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SkillGateWay interface {
	GetSkills() []Skill
}
