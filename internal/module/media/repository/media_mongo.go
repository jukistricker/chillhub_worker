package repository

import (
	"context"
	"time"
	"chillhub/internal/module/media/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"chillhub/internal/shared/catalog"
)

type MediaRepository interface {
	Insert(ctx context.Context, media *model.Media) error
	UpdateStatus(ctx context.Context, id string, status model.MediaStatus) error
	FindByID(ctx context.Context, objID primitive.ObjectID) (*model.Media, error)
}

type mongoRepo struct {
	col *mongo.Collection
}

// UpdateStatus: Hiện thực hóa logic để cập nhật trạng thái video
func (r *mongoRepo) UpdateStatus(ctx context.Context, id string, status model.MediaStatus) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return catalog.BadRequest.Err(err, "media.invalid_id")
	}

	// Thực hiện Update trong MongoDB
	_, err = r.col.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"status":     status,
				"updated_at": time.Now().Unix(),
			},
		},
	)

	return nil
}

// Cài đặt FindByID cho struct mongoRepo (KHÔNG dùng MediaRepository làm receiver)
func (r *mongoRepo) FindByID(ctx context.Context, objID primitive.ObjectID) (*model.Media, error) {
	var m model.Media
	err := r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&m)
	if err != nil {
		return nil, catalog.NotFound.Err(
			err,
			"media.not_found",
		)
	}
	return &m, nil
}

func NewMediaMongo(db *mongo.Database) MediaRepository {
	return &mongoRepo{
		col: db.Collection("media"),
	}
}

func (r *mongoRepo) Insert(ctx context.Context, media *model.Media) error {
	media.CreatedAt = time.Now().Unix()
	_, err := r.col.InsertOne(ctx, media)
	return err
}
