package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wFercho/mines_microservice/internal/domain/mine_nodes3d"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MineNodes3DMongoRepository struct {
	collection *mongo.Collection
}

func NewMineNodes3DMongoRepository(db *mongo.Database) *MineNodes3DMongoRepository {
	return &MineNodes3DMongoRepository{
		collection: db.Collection("mine_nodes3d"),
	}
}

func (r *MineNodes3DMongoRepository) Create(mn *mine_nodes3d.MineNodes3D) (*mine_nodes3d.MineNodes3D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	model := FromDomainToMineNodes3dMongoModel(mn)
	_, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	return mn, nil
}

func (r *MineNodes3DMongoRepository) FindByID(id uuid.UUID) (*mine_nodes3d.MineNodes3D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id.String())
	if err != nil {
		return nil, errors.New("ID no válido")
	}

	filter := bson.M{"_id": objectID}
	var model MineNodes3DMongoModel

	err = r.collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *MineNodes3DMongoRepository) FindByMineID(mineID uuid.UUID) (*mine_nodes3d.MineNodes3D, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"mine_id": mineID.String()}
	var model MineNodes3DMongoModel

	err := r.collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}

func (r *MineNodes3DMongoRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id.String())
	if err != nil {
		return errors.New("ID no válido")
	}

	filter := bson.M{"_id": objectID}
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("registro no encontrado")
	}

	return nil
}

func (r *MineNodes3DMongoRepository) DeleteByMineID(mineID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"mine_id": mineID.String()}
	_, err := r.collection.DeleteMany(ctx, filter)
	return err
}
