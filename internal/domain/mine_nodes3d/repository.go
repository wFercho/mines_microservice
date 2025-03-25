package mine_nodes3d

import (
	"github.com/google/uuid"
)

type MineNodes3DRepository interface {
	Create(mn *MineNodes3D) (*MineNodes3D, error)
	FindByID(id uuid.UUID) (*MineNodes3D, error)
	FindByMineID(mine_id uuid.UUID) (*MineNodes3D, error)
	Delete(id uuid.UUID) error
	DeleteByMineID(mine_id uuid.UUID) error
}
