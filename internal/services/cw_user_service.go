package services

import (
	"cmd/internal/db"
)

type CwUserService struct {
	repository *db.Repository
}

func NewCwUserService(repository *db.Repository) *CwUserService {
	return &CwUserService{repository: repository}
}
