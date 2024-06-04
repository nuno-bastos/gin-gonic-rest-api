package interfaces

import (
	"context"
	"golang-gin-api/model"
)

type TagRepository interface {
	FindAll(ctx context.Context) ([]model.Tags, error)
	FindById(ctx context.Context, id uint) (model.Tags, error)
	Save(ctx context.Context, tag model.Tags) (model.Tags, error)
	Delete(ctx context.Context, tag model.Tags) error
}
