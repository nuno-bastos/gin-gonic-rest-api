package interfaces

import (
	"context"

	model "golang-gin-api/model"
)

type TagService interface {
	FindAll(ctx context.Context) ([]model.Tags, error)
	FindByID(ctx context.Context, id uint) (model.Tags, error)
	Save(ctx context.Context, tag model.Tags) (model.Tags, error)
	Delete(ctx context.Context, tag model.Tags) error
}
