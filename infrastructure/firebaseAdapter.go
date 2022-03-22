package infrastructure

import (
	"personal-page-back/domain"
)

type firebaseAdapter struct {
	baseUrl string
}

func NewFirebaseAdapter(url string) *firebaseAdapter {
	return &firebaseAdapter{}
}

func (f *firebaseAdapter) GetSkills() []domain.Skill {
	return []domain.Skill{
		{
			Name:  "Devops",
			Count: 0,
		}, {
			Name:  "Extra",
			Count: 0,
		},
		{
			Name:  "Elixir",
			Count: 0,
		},
		{
			Name:  "Excel",
			Count: 0,
		},
	}
}
