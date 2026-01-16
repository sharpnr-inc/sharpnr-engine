package org_signup_repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type orgMainSignUpRepo struct {
	db *gorm.DB
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
