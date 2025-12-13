package db

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Create(entity).Error
}

func (r *Repository[T]) Update(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Save(entity).Error
}

func (r *Repository[T]) Delete(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Delete(entity).Error
}

func (r *Repository[T]) CountByID(ctx context.Context, db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindByID(ctx context.Context, db *gorm.DB, entity *T, id string) error {
	return db.WithContext(ctx).Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) FindAll(ctx context.Context, db *gorm.DB) ([]T, error) {
	var entities []T
	err := db.WithContext(ctx).Find(&entities).Error
	return entities, err
}
