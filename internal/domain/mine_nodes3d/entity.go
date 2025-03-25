package mine_nodes3d

import "github.com/google/uuid"

type Node3D struct {
	ID          string
	Zone        Zone
	Connections []string
	Position    Position
	Color       string
	Sensors     []Sensor
}

type Zone struct {
	Category string
	Name     string
}

type Position struct {
	X int
	Y int
	Z int
}

type Sensor struct {
	Category string
	Unit     string
}

type MineNodes3D struct {
	ID     uuid.UUID
	MineId uuid.UUID
	Nodes  []Node3D
}

func NewMineNodes3D(mineId uuid.UUID, nodes []Node3D) (*MineNodes3D, error) {
	return &MineNodes3D{
		ID:     uuid.New(),
		MineId: mineId,
		Nodes:  nodes,
	}, nil
}
