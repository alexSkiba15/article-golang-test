package adapters

import (
	"context"
	"errors"
	"rest-project/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Model interface {
		entities.Article | entities.Base
	}

	IModel[TE Model] interface {
		ToEntity() *TE
		FromEntity(entity *TE) any
	}

	Repo[T Model] interface {
		Create(entity *T, ctx context.Context) error
		BulkCreate(entity *[]T, ctx context.Context) error
		GetById(id uuid.UUID, ctx context.Context) (*T, error)
		Get(params *T, ctx context.Context) (*T, error)
		GetAll(ctx context.Context) (*[]T, error)
		Update(entity *T, ctx context.Context) error
		UpdateAll(entity *[]T, ctx context.Context) error
		DeleteById(id uuid.UUID, ctx context.Context) error
	}

	ArticleRepository interface {
		Repo[entities.Article]
	}

	Repository[TM IModel[TE], TE Model] struct {
		db *gorm.DB
	}
)

func NewRepository[TM IModel[TE], TE Model](db *gorm.DB) *Repository[TM, TE] {
	return &Repository[TM, TE]{
		db: db,
	}
}

func (r *Repository[TM, TE]) Create(entity *TE, ctx context.Context) error {
	var model TM
	model = model.FromEntity(entity).(TM)
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *Repository[TM, TE]) BulkCreate(entity *[]TE, ctx context.Context) error {
	var models []TM
	for _, e := range *entity {
		var model TM
		models = append(models, model.FromEntity(&e).(TM))
	}
	return r.db.WithContext(ctx).Create(&models).Error
}

func (r *Repository[TM, TE]) GetById(id uuid.UUID, ctx context.Context) (*TE, error) {
	var model TM
	err := r.db.WithContext(ctx).Model(&model).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *Repository[TM, TE]) Get(params *TE, ctx context.Context) (*TE, error) {
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

func (r *Repository[TM, TE]) Where(params *TE, ctx context.Context) (*[]TE, error) {
	var entitiesObjects []TE
	err := r.db.WithContext(ctx).Where(&params).Find(&entitiesObjects).Error
	if err != nil {
		return nil, err
	}
	return &entitiesObjects, nil
}

func (r *Repository[TM, TE]) Update(entity *TE, ctx context.Context) error {
	var model TM
	_ = model.FromEntity(entity).(TM)
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *Repository[TM, TE]) UpdateAll(entities *[]TE, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *Repository[TM, TE]) DeleteById(id uuid.UUID, ctx context.Context) error {
	var entity TE
	err := r.db.WithContext(ctx).Delete(&entity, id).Error
	return err
}

func (r *Repository[TM, TE]) SkipTake(skip int, take int, ctx context.Context) (*[]TE, error) {
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
