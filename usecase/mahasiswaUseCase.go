package usecase

import (
	"enigma-mhs/models"
	"enigma-mhs/repository"
)

type IMahasiswaUseCase interface {
	GetMahasiswaById(Id string) (*models.Mahasiswa,error)
}

type MahasiswaUseCase struct {
	repo repository.IMahasiswaRepository
}

func NewMahasiswaUseCase(repo repository.IMahasiswaRepository) IMahasiswaUseCase {
	return &MahasiswaUseCase{repo}
}

func (m *MahasiswaUseCase) GetMahasiswaById(Id string) (*models.Mahasiswa,error) {
	return m.repo.SelectById(Id)
}