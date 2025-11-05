package utils

type PaginationResult[T any] struct {
	Page       int
	Limit      int
	Total      int64
	TotalPages int
	Data       []T
}

// func Paginate[T any](c *gin.Context, db *gorm.DB) (*PaginationResult[T], error) {

// }
