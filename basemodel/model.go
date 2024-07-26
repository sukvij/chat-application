package basemodel

import (
	"time"
)

type BaseModel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// DeletedAt *time.Time
}

// func (m *BaseModel) BeforeInsert(db gorm.DB) error {
// 	now := time.Now()
// 	if m.CreatedAt.IsZero() {
// 		m.CreatedAt = now
// 	}
// 	if m.UpdatedAt.IsZero() {
// 		m.UpdatedAt = now
// 	}
// 	return nil
// }

// func (m *BaseModel) BeforeUpdate(db gorm.DB) error {
// 	m.UpdatedAt = time.Now()
// 	return nil
// }
