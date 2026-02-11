package sero

import (
	"context"
	"gorm.io/gorm"
)

type OutboxRepository interface {
	Save(ctx context.Context, tx *gorm.DB, event *SeroOutbox) error
	FetchPending(ctx context.Context, batchSize int) ([]SeroOutbox, error)
	UpdateStatus(ctx context.Context, id uint, status string, errStr string) error
}

type GormOutboxRepo struct {
	db *gorm.DB
}

func NewGormOutboxRepo(db *gorm.DB) *GormOutboxRepo {
	return &GormOutboxRepo{db: db}
}

func (r *GormOutboxRepo) Save(ctx context.Context, tx *gorm.DB, event *SeroOutbox) error {
	return tx.WithContext(ctx).Create(event).Error
}

func (r *GormOutboxRepo) FetchPending(ctx context.Context, batchSize int) ([]SeroOutbox, error) {
	var events []SeroOutbox
	err := r.db.WithContext(ctx).
		Where("status = ? AND retries < ?", "PENDING", 3).
		Order("created_at asc").
		Limit(batchSize).
		Find(&events).Error
	return events, err
}

func (r *GormOutboxRepo) UpdateStatus(ctx context.Context, id uint, status string, errStr string) error {
	updates := map[string]interface{}{"status": status, "last_error": errStr}
	if status == "FAILED" {
		updates["retries"] = gorm.Expr("retries + ?", 1)
	}
	return r.db.WithContext(ctx).Model(&SeroOutbox{}).Where("id = ?", id).Updates(updates).Error
}
