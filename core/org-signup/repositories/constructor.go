package org_signup_repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type orgSignUpRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewOrgSignUpRepository(db *gorm.DB, redis *redis.Client) *orgSignUpRepo {
	return &orgSignUpRepo{
		db:    db,
		redis: redis,
	}
}
