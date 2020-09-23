package usecase

import (
	"enigma-mhs/models"
	"enigma-mhs/repository"
)

type IMhsMatkulUseCase interface {
	GetAllMhsById(MhsId string) ([]*models.MhsMatkul,error)
}

type MhsMatkulUseCase struct {
	repo repository.IMhsMatkulRepository
}

func NewMhsMatkulUseCase(repo repository.IMhsMatkulRepository) IMhsMatkulUseCase {
	return &MhsMatkulUseCase{repo}
}

func (b *MhsMatkulUseCase) GetAllMhsById(MhsId string) ([]*models.MhsMatkul,error) {
	return b.repo.SelectByMHSId(MhsId)
}