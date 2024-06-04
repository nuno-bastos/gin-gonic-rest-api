package service

import (
	"context"

	model "golang-gin-api/model"
	interfaces "golang-gin-api/repo/interface"
	services "golang-gin-api/service/interface"
)

type tagService struct {
	tagRepo interfaces.TagRepository
}

func NewTagService(repo interfaces.TagRepository) services.TagService {
	return &tagService{
		tagRepo: repo,
	}
}

func (p *tagService) FindAll(ctx context.Context) ([]model.Tags, error) {
	tags, err := p.tagRepo.FindAll(ctx)
	return tags, err
}

func (p *tagService) FindByID(ctx context.Context, id uint) (model.Tags, error) {
	tag, err := p.tagRepo.FindById(ctx, id)
	return tag, err
}

func (p *tagService) Save(ctx context.Context, tag model.Tags) (model.Tags, error) {
	tag, err := p.tagRepo.Save(ctx, tag)

	return tag, err
}

func (p *tagService) Delete(ctx context.Context, tag model.Tags) error {
	err := p.tagRepo.Delete(ctx, tag)

	return err
}
