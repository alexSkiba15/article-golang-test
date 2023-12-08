package adapters

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rest-project/src/domain/entities"
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

	UnitOfWork interface {
		ArticleRepository() ArticleRepository
		Do(func(UnitOfWork) error) error
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
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[TM, TE]) BulkCreate(entity *[]TE, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[TM, TE]) GetById(id uuid.UUID, ctx context.Context) (*TE, error) {
	var entity TE
	err := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *Repository[TM, TE]) Get(params *TE, ctx context.Context) (*TE, error) {
	var entity TE
	r.db.WithContext(ctx).Where(&params).First(&entity)
	if entity == nil {
		notFound := NotFoundError{Name: "Article"}
		return nil, notFound.Error()
	}
	return &entity, nil
}

func (r *Repository[TM, TE]) GetAll(ctx context.Context) (*[]TE, error) {
	var entities_ []TE
	err := r.db.WithContext(ctx).Find(&entities_).Error
	if err != nil {
		return nil, err
	}
	return &entities_, nil
}

func (r *Repository[TM, TE]) Where(params *TE, ctx context.Context) (*[]TE, error) {
	var entities_ []TE
	err := r.db.WithContext(ctx).Where(&params).Find(&entities_).Error
	if err != nil {
		return nil, err
	}
	return &entities_, nil
}

func (r *Repository[TM, TE]) Update(entity *TE, ctx context.Context) error {
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
	var entities_ []TE
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities_).Error
	if err != nil {
		return nil, err
	}
	return &entities_, nil
}

func (r *Repository[TM, TE]) Count(ctx context.Context) int64 {
	var entity TE
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}
