package restaurantrepo

// func (r *restaurantRepo) WithTransaction(ctx context.Context, fn func(tx RestaurantRepository) error) error {
// 	tx := r.db.Begin()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	txRepo := &restaurantRepo{db: tx}

// 	if err := fn(txRepo); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return tx.Commit().Error
// }
