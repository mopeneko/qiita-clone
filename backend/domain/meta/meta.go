package meta

import "time"

// Meta is common meta info
type Meta struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
