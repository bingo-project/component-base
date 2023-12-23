package gormutil

// DefaultLimit define the default number of records to be retrieved.
const DefaultLimit = 15

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
