package pagination

import (
	"math"

	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, pagination *Pagination, model any, args ...any) func(db *gorm.DB) *gorm.DB {
	var query any
	var rest []any

	if len(args) > 0 {
		query = args[0]
	}

	if len(args) > 1 {
		rest = args[1:]
	}

	var totalRows int64

	db.Model(model).Where(query, rest...).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("id asc")
	}
}
