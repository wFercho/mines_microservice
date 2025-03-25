package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wFercho/mines_microservice/internal/domain/mine_nodes3d"
)

type MineNodes3DUseCase struct {
	Repository mine_nodes3d.MineNodes3DRepository
}

func NewMineNodes3DUseCase(repo mine_nodes3d.MineNodes3DRepository) *MineNodes3DUseCase {
	return &MineNodes3DUseCase{Repository: repo}
}

func (uc *MineNodes3DUseCase) Create(mineNodes3D *mine_nodes3d.MineNodes3D) (*mine_nodes3d.MineNodes3D, error) {
	if mineNodes3D == nil {
		return nil, errors.New("los datos de MineNodes3D no pueden ser nulos")
	}

	createdMineNodes3D, err := uc.Repository.Create(mineNodes3D)
	if err != nil {
		return nil, err
	}

	return createdMineNodes3D, nil
}

func (uc *MineNodes3DUseCase) FindByID(id uuid.UUID) (*mine_nodes3d.MineNodes3D, error) {
	mineNodes3D, err := uc.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if mineNodes3D == nil {
		return nil, errors.New("registro no encontrado")
	}

	return mineNodes3D, nil
}

func (uc *MineNodes3DUseCase) FindByMineID(mineID uuid.UUID) (*mine_nodes3d.MineNodes3D, error) {
	mineNodes3D, err := uc.Repository.FindByMineID(mineID)
	if err != nil {
		return nil, err
	}

	if mineNodes3D == nil {
		return nil, errors.New("registro no encontrado")
	}

	return mineNodes3D, nil
}

func (uc *MineNodes3DUseCase) Delete(id uuid.UUID) error {
	return uc.Repository.Delete(id)
}

func (uc *MineNodes3DUseCase) DeleteByMineID(mineID uuid.UUID) error {
	return uc.Repository.DeleteByMineID(mineID)
}
