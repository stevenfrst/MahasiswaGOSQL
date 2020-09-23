package usecase

import (
	"enigma-mhs/models"
	"enigma-mhs/repository"
)

type IMatkulUseCase interface {
	GetMatkulById(Id string) (*models.Matkul,error)
}

type MatkulUseCase struct {
	repo  repository.IMatkulRepository
}

func NewMatkulUseCase(repo repository.IMatkulRepository) IMatkulUseCase {
	return &MatkulUseCase{repo}
}

func (x *MatkulUseCase) GetMatkulById(Id string) (*models.Matkul,error) {
	return x.repo.SelectById(Id)
}