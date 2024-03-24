package adapters

import (
	"context"
	"errors"

	"github.com/alexSkiba15/article-golang-test/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Model interface {
		entities.Article | entities.Base
	}

	IModel[TE Model, TM any] interface {
		ToEntity() *TE
		FromEntity(entity *TE) TM
	}

	Repo[T Model] interface {
		Create(ctx context.Context, entity *T) error
		BulkCreate(ctx context.Context, entity *[]T) error
		GetById(ctx context.Context, id uuid.UUID) (*T, error)
		Get(ctx context.Context, params *T) (*T, error)
		GetAll(ctx context.Context) (*[]T, error)
		Update(ctx context.Context, entity *T) error
		UpdateAll(ctx context.Context, entity *[]T) error
		DeleteById(ctx context.Context, id uuid.UUID) error
	}

	ArticleRepository interface {
		Repo[entities.Article]
	}

	Repository[TM IModel[TE, TM], TE Model] struct {
		db *gorm.DB
	}
)

func NewRepository[TM IModel[TE, TM], TE Model](db *gorm.DB) *Repository[TM, TE] {
	return &Repository[TM, TE]{
		db: db,
	}
}

func (r *Repository[TM, TE]) Create(ctx context.Context, entity *TE) error {
	var model TM
	model = model.FromEntity(entity)
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *Repository[TM, TE]) BulkCreate(ctx context.Context, entity *[]TE) error {
	var models []TM
	for _, e := range *entity {
		var model TM
		models = append(models, model.FromEntity(&e))
	}
	return r.db.WithContext(ctx).Create(&models).Error
}

func (r *Repository[TM, TE]) GetById(ctx context.Context, id uuid.UUID) (*TE, error) {
	var model TM
	err := r.db.WithContext(ctx).Model(&model).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *Repository[TM, TE]) Get(ctx context.Context, params *TE) (*TE, error) {
	var model TM
	result := r.db.WithContext(ctx).Where(&params).First(&model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		notFound := NotFoundError{Name: "Article"}
		return nil, notFound.Error()
	}
	return model.ToEntity(), nil
}

func (r *Repository[TM, TE]) GetAll(ctx context.Context) (*[]TE, error) {
	var model TM
	var entitiesObjects []TE
	err := r.db.WithContext(ctx).Model(model).Find(&entitiesObjects).Error
	if err != nil {
		return nil, err
	}
	return &entitiesObjects, nil
}

func (r *Repository[TM, TE]) Where(ctx context.Context, params *TE) (*[]TE, error) {
	var entitiesObjects []TE
	err := r.db.WithContext(ctx).Where(&params).Find(&entitiesObjects).Error
	if err != nil {
		return nil, err
	}
	return &entitiesObjects, nil
}

func (r *Repository[TM, TE]) Update(ctx context.Context, entity *TE) error {
	var model TM
	_ = model.FromEntity(entity)
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *Repository[TM, TE]) UpdateAll(ctx context.Context, entities *[]TE) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *Repository[TM, TE]) DeleteById(ctx context.Context, id uuid.UUID) error {
	var entity TE
	err := r.db.WithContext(ctx).Delete(&entity, id).Error
	return err
}

func (r *Repository[TM, TE]) SkipTake(ctx context.Context, skip int, take int) (*[]TE, error) {
	var objects []TE
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&objects).Error
	if err != nil {
		return nil, err
	}
	return &objects, nil
}

func (r *Repository[TM, TE]) Count(ctx context.Context) int64 {
	var entity TE
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}
