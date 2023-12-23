package gormutil

// DefaultLimit define the default number of records to be retrieved.
const DefaultLimit = 15

// LimitAndOffset contains offset and limit fields.
type LimitAndOffset struct {
	Offset int
	Limit  int
}

type ListRequest struct {
	// Page
	Page int `form:"page"`

	// Offset
	Offset int `form:"offset"`

	// Limit
	Limit int `form:"limit"`

	// Order by field.
	Order string `form:"order"`

	// Sort: asc or desc.
	Sort string `form:"sort"`
}

type ListResponse struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}

// SetDefaultParams Set default params if not exist.
func (req *ListRequest) SetDefaultParams() {
	// Sort
	if req.Sort == "" {
		req.Sort = "id"
	}

	// Order
	if req.Order == "" {
		req.Order = "desc"
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
