package repository

import (
	"context"

	"github.com/gsy/store/internal/domain"
	"gorm.io/gorm"
)

type Model interface{}

type Convertor[M Model, E domain.Entity] interface {
	FromEntity(entity *E) (model *M)
	ToEntity(model *M) (entity *E)
}

func NewRepository[M Model, E domain.Entity](db *gorm.DB, conv Convertor[M, E]) *GeneralRepository[M, E] {
	return &GeneralRepository[M, E]{
		db:   db,
		conv: conv,
	}
}

type GeneralRepository[M Model, E domain.Entity] struct {
	db   *gorm.DB
	conv Convertor[M, E]
}

func (repo *GeneralRepository[M, E]) Create(ctx context.Context, entity *E) (result *E, err error) {
	model := repo.conv.FromEntity(entity)
	if err = repo.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}
	return repo.conv.ToEntity(model), nil
}
