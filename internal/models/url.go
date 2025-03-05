package models

import (
	"time"
)

type Url struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Hash      string    `json:"hash" gorm:"primarykey;uniqueIndex:idx_hash"`
	Permalink string    `json:"permalink"`
}
