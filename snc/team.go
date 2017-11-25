package snc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// todo use deleted at fields
//--------------------------------------------------------------------------------------------------------------------//
// Team
//--------------------------------------------------------------------------------------------------------------------//
type Team struct {
	ID         uint       `gorm:"primary_key"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	Name       string     `json:"name" gorm:"not null;unique_index"`
	Division   Division   `json:"division"`
	DivisionID uint       `json:"-"`
}

func CreateTeam(t Team, DB *gorm.DB) error {
	DB.Create(&t)
	return DB.Error
}

func FetchTeam(id uint, DB *gorm.DB) (Team, error) {
	fmt.Println("Fetching team")
	var team Team
	DB.Preload("Division").Where("deleted_at IS NULL").First(&team, id)
	return team, DB.Error
}

func FetchTeams(DB *gorm.DB) ([]Team, error) {
	teams := make([]Team, 0)
	DB = DB.Preload("Division").Where("deleted_at IS NULL").Find(&teams)
	return teams, DB.Error
}

func TeamCalled(name string, DB *gorm.DB) (Team, error) {
	var team Team
	DB.Preload("Division").Where("name = ? AND deleted_at IS NULL", name).First(&team)
	return team, DB.Error
}

func UpdateTeam(t Team, DB *gorm.DB) error {
	DB.Where("ID = ? AND deleted_at IS NULL", t.ID).Save(&t)
	return DB.Error
}

func DeleteTeam(id uint, DB *gorm.DB) error {
	var team Team
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&team)
	return DB.Error
}

//--------------------------------------------------------------------------------------------------------------------//
// Division
//--------------------------------------------------------------------------------------------------------------------//
type Division struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name" gorm:"not null;unique_index"`
}

func CreateDivision(d Division, DB *gorm.DB) error {
	DB.Create(&d)
	return DB.Error
}

func FetchDivision(id uint, DB *gorm.DB) (Division, error) {
	var d Division
	DB.Where("ID = ? AND deleted_at IS NULL", id).First(&d)
	return d, DB.Error
}

func FetchDivisions(DB *gorm.DB) ([]Division, error) {
	d := make([]Division, 0)
	DB.Where("deleted_at IS NULL").Find(&d)
	return d, DB.Error
}

func UpdateDivision(d Division, DB *gorm.DB) error {
	DB.Where("ID = ? AND deleted_at IS NULL", d.ID).Save(&d)
	return DB.Error
}

func DeleteDivision(id int, DB *gorm.DB) error {
	var div Division
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&div)
	return DB.Error
}

//--------------------------------------------------------------------------------------------------------------------//
// Player (Goalies included)
//--------------------------------------------------------------------------------------------------------------------//
type Player struct {
	gorm.Model
	JerseyNumber uint   `json:"jerseyNumber"`
	Name         string `json:"name"`
	Position     string `json:"position"`
}

func CreatePlayer(p Player, DB *gorm.DB) error {
	DB.Create(&p)
	return DB.Error
}

func FetchPlayer(id uint, DB *gorm.DB) (Player, error) {
	var p Player
	DB.Where("ID = ? AND deleted_at IS NULL", id).First(&p)
	return p, DB.Error
}

func FetchPlayers(DB *gorm.DB) ([]Player, error) {
	p := make([]Player, 0)
	DB.Where("deleted_at IS NULL").Find(&p)
	return p, DB.Error
}

func UpdatePlayer(p Player, DB *gorm.DB) error {
	DB.Where("ID = ? AND deleted_at IS NULL", p.ID).Save(&p)
	return DB.Error
}

func DeletePlayer(id int, DB *gorm.DB) error {
	var p Player
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&p)
	return DB.Error
}
