package snc

import (
	"github.com/jinzhu/gorm"
)

const (
	Avondale = "Avondale"
	Botany   = "Botany"
)

type Rink struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;unique_index"`
}

func CreateRink(r Rink, DB *gorm.DB) error {
	DB.Create(&r)
	return DB.Error
}

func FetchRink(id uint, DB *gorm.DB) (Rink, error) {
	var r Rink
	DB.Where("ID = ? AND deleted_at IS NULL", id).First(&r)
	return r, DB.Error
}

func FetchRinks(DB *gorm.DB) ([]Rink, error) {
	r := make([]Rink, 0)
	DB.Where("deleted_at IS NULL").Find(&r)
	return r, DB.Error
}

func UpdateRink(r Rink, DB *gorm.DB) error {
	DB.Where("ID = ? AND deleted_at IS NULL", r.ID).Save(&r)
	return DB.Error
}

func DeleteRink(id int, DB *gorm.DB) error {
	var r Rink
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&r)
	return DB.Error
}
