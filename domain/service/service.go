package service

import (
	"github.com/bariasabda/test-flash-coffee/domain/repository"
	"github.com/bariasabda/test-flash-coffee/domain/service/recipe"
)

type Service struct {
	Recipe recipe.RecipeSvc
}

func New(repo repository.Repository) Service {
	return Service{
		Recipe: recipe.New(repo.Recipe, repo.User),
	}
}
