package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Sample struct {
	Id   int `gorm:"AUTO_INCREMENT"`
	Name string
}

func FindSampleAll(db *gorm.DB) ([]Sample, error) {

	sample := []Sample{}
	if db.Table("samples").Find(&sample).Error != nil {
		return nil, fmt.Errorf("error")
	}

	return sample, nil
}
