package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rakhiazfa/fiber-boilerplate/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	*Repository[entity.User, uuid.UUID]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository[entity.User, uuid.UUID]{db},
	}
}

func (r *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return NewUserRepository(tx)
}

func (r *UserRepository) WithContext(ctx context.Context) *UserRepository {
	return NewUserRepository(r.db.WithContext(ctx))
}
