package gormutil

import "strings"

const (
	// DefaultLimit define the default number of records to be retrieved.
	DefaultLimit = 15

	ASC  = "ASC"
	DESC = "DESC"
)

// ListOptions contains offset and limit fields.
type ListOptions struct {
	// Page
	Page int `form:"page"`

	// Offset
	Offset int `form:"offset"`

	// Limit
	Limit int `form:"limit"`

	// Order asc or desc.
	Order string `form:"order"`

	// Sort field.
	Sort string `form:"sort"`
}

// SetDefaultOptions Set default options if not exist.
func (req *ListOptions) SetDefaultOptions() {
	// Sort
	if req.Sort == "" {
		req.Sort = "id"
	}

	// Order
	req.Order = strings.ToUpper(req.Order)
	if req.Order != ASC {
		req.Order = DESC
	}

	// Page size
	if req.Limit <= 0 {
		req.Limit = DefaultLimit
	}

	// Page
	if req.Page <= 0 {
		req.Page = 1
	}

	// Offset
	if req.Offset == 0 && req.Page > 0 {
		req.Offset = (req.Page - 1) * req.Limit
	}
}
