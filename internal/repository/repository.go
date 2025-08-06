package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repository[T any, ID any] struct {
	db *gorm.DB
}

func (r *Repository[T, ID]) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository[T, ID]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T, ID]) FindById(id ID) (*T, error) {
	var entity T

	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &entity, nil
}

func (r *Repository[T, ID]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T, ID]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *Repository[T, ID]) CountBy(field string, value any, excludedIds ...ID) (int64, error) {
	var count int64

	q := r.db.Model(new(T)).Where(fmt.Sprintf("%s = ?", field), value)
	if len(excludedIds) > 0 {
		q = q.Where("id NOT IN ?", excludedIds)
	}

	err := q.Count(&count).Error

	return count, err
}
