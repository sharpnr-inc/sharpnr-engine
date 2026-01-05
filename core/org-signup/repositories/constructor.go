package org_signup_repositories

import (
	"context"
	"sharpnr-engine/core/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type orgMainSignUpRepo struct {
	db *gorm.DB
}

// FindOrgBySharpnrID implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) FindOrgBySharpnrID(ctx context.Context, Id string) (models.OrganizationSignUpStaging, bool, error) {
	panic("unimplemented")
}

func NewOrgMainSignUpRepository(db *gorm.DB) *orgMainSignUpRepo {
	return &orgMainSignUpRepo{
		db: db,
	}
}

type orgTempSignUpRepo struct {
	redis *redis.Client
}

func NewTempOrgSignUpRepository(redis *redis.Client) *orgTempSignUpRepo {
	return &orgTempSignUpRepo{
		redis: redis,
	}
}
