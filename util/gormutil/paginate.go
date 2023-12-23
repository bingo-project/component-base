package gormutil

import (
	"gorm.io/gorm"
)

// Paginate the query.
func Paginate(db *gorm.DB, req *ListOptions, data any) (count int64, err error) {
	// Set default params
	req.SetDefaultOptions()

	// Query
	err = db.Order(req.Sort + " " + req.Order).
		Offset(req.Offset).
		Limit(req.Limit).
		Find(data).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	return
}

// Paginator returns a gorm scope paginator.
func Paginator(req *ListOptions) func(db *gorm.DB) *gorm.DB {
	// Set default params
	req.SetDefaultOptions()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(req.Offset).Limit(req.Limit)
	}
}
