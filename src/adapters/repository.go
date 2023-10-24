package adapters

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model interface {
	GetId() uuid.UUID
}

type Repo[T Model] interface {
	Create(entity *T, ctx context.Context) error
	BulkCreate(entity *[]T, ctx context.Context) error
	GetById(id uuid.UUID, ctx context.Context) (*T, error)
	Get(params *T, ctx context.Context) (*T, error)
	GetAll(ctx context.Context) (*[]T, error)
	Update(entity *T, ctx context.Context) error
	UpdateAll(entity *[]T, ctx context.Context) error
	DeleteById(id uuid.UUID, ctx context.Context) error
}

type Repository[T Model] struct {
	db *gorm.DB
}

func NewRepository[T Model](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) Create(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[T]) BulkCreate(entity *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository[T]) GetById(id uuid.UUID, ctx context.Context) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *Repository[T]) Get(params *T, ctx context.Context) (*T, error) {
	var entity T
	r.db.WithContext(ctx).Where(&params).First(&entity)
	if entity == nil {
		notFound := NotFoundError{Name: "Article"}
		return nil, notFound.Error()
	}
	return &entity, nil
}

func (r *Repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Where(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Where(&params).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Update(entity *T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *Repository[T]) UpdateAll(entities *[]T, ctx context.Context) error {
	return r.db.WithContext(ctx).Save(&entities).Error
}

func (r *Repository[T]) DeleteById(id uuid.UUID, ctx context.Context) error {
	var entity T
	err := r.db.WithContext(ctx).Delete(&entity, id).Error
	return err
}

func (r *Repository[T]) SkipTake(skip int, take int, ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Offset(skip).Limit(take).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *Repository[T]) Count(ctx context.Context) int64 {
	var entity T
	var count int64
	r.db.WithContext(ctx).Model(&entity).Count(&count)
	return count
}
