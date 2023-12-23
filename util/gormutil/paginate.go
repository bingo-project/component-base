package gormutil

import (
	"gorm.io/gorm"
)

// Paginate the query.
func Paginate(db *gorm.DB, req *ListRequest, data any) (res *ListResponse, err error) {
	// Set default params
	req.SetDefaultParams()

	var count int64

	// Query
	err = db.Order(req.Sort + " " + req.Order).
		Offset(req.Offset).
		Limit(req.Limit).
		Find(data).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
	if err != nil {
		return nil, err
	}

	return &ListResponse{Total: count, Data: data}, nil
}

// Paginator returns a gorm scope paginator.
func Paginator(req *ListRequest) func(db *gorm.DB) *gorm.DB {
	// Set default params
	req.SetDefaultParams()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(req.Offset).Limit(req.Limit)
	}
}
