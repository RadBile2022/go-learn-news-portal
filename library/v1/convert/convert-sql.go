package convert

import (
	"database/sql"
	"time"
)

func StringSql(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func Float64Sql(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0.0
}

func Int64Sql(ni sql.NullInt64) int64 {
	if ni.Valid {
		return ni.Int64
	}
	return 0
}

func TimeSql(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
