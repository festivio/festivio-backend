package service

import "github.com/festivio/festivio-backend/domain"

func (s service) GetTeam() (*domain.Team, error) {
	var team domain.Team
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Role == "Аниматор" {
			team.Count.Animators++
		} else if user.Role == "Менеджер" {
			team.Count.Managers++
		} else if user.Role == "Основатель" {
			team.Count.Founders++
		}
		team.Count.TotalCount++
	}
	team.Employees = users

	return &team, nil
}