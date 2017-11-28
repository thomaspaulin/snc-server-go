package snc

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	Avondale = "Avondale"
	Botany   = "Botany"
)

type Rink struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name" gorm:"not null;unique_index;primary_key"`
}

func CreateRink(r Rink, DB *gorm.DB) error {
	rink, _ := FetchRinkNamed(r.Name, DB)
	if rink.ID != 0 {
		r.ID = rink.ID
		return UpdateRink(r, DB)
	}
	res := DB.Create(&r)
	return res.Error
}

func FetchRink(id uint, DB *gorm.DB) (Rink, error) {
	var r Rink
	res := DB.Where("id = ? AND deleted_at IS NULL", id).First(&r)
	return r, res.Error
}

func FetchRinkNamed(name string, DB *gorm.DB) (Rink, error) {
	var r Rink
	res := DB.Where("Name = ? AND deleted_at IS NULL", name).First(&r)
	return r, res.Error
}

func FetchRinks(DB *gorm.DB) ([]Rink, error) {
	r := make([]Rink, 0)
	res := DB.Where("deleted_at IS NULL").Find(&r)
	return r, res.Error
}

func UpdateRink(r Rink, DB *gorm.DB) error {
	res := DB.Where("id = ? AND deleted_at IS NULL", r.ID).Save(&r)
	return res.Error
}

func DeleteRink(id int, DB *gorm.DB) error {
	var r Rink
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&r)
	return res.Error
}
