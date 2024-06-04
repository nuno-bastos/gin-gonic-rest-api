package repo

import (
	"context"

	model "golang-gin-api/model"
	interfaces "golang-gin-api/repo/interface"

	"gorm.io/gorm"
)

type tagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(DB *gorm.DB) interfaces.TagRepository {
	return &tagRepository{DB}
}

func (p *tagRepository) FindAll(ctx context.Context) ([]model.Tags, error) {
	var tags []model.Tags
	err := p.DB.Find(&tags).Error

	return tags, err
}

func (p *tagRepository) FindById(ctx context.Context, id uint) (model.Tags, error) {
	var tag model.Tags
	err := p.DB.First(&tag, id).Error

	return tag, err
}

func (p *tagRepository) Save(ctx context.Context, tag model.Tags) (model.Tags, error) {
	err := p.DB.Save(&tag).Error

	return tag, err
}

func (p *tagRepository) Delete(ctx context.Context, tag model.Tags) error {
	err := p.DB.Delete(&tag).Error

	return err
}
