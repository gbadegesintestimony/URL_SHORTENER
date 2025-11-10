package models

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	Short     string `gorm:"uniqueIndex;not null"`
	Original  string `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}
