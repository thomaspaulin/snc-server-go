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
	ID           uint       `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
	Name         string     `json:"name" gorm:"not null;unique_index"`
	DivisionName string     `json:"divisionName"`
}

func CreateTeam(t Team, DB *gorm.DB) error {
	res := DB.Create(&t)
	return res.Error
}

func FetchTeam(id uint, DB *gorm.DB) (Team, error) {
	fmt.Println("Fetching team")
	var team Team
	res := DB.Where("deleted_at IS NULL").First(&team, id)
	return team, res.Error
}

func FetchTeams(DB *gorm.DB) ([]Team, error) {
	teams := make([]Team, 0)
	res := DB.Where("deleted_at IS NULL").Find(&teams)
	return teams, res.Error
}

func TeamCalled(name string, DB *gorm.DB) (Team, error) {
	var team Team
	res := DB.Where("name = ? AND deleted_at IS NULL", name).First(&team)
	return team, res.Error
}

func UpdateTeam(t Team, DB *gorm.DB) error {
	res := DB.Where("ID = ? AND deleted_at IS NULL", t.ID).Save(&t)
	return res.Error
}

func DeleteTeam(id uint, DB *gorm.DB) error {
	var team Team
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&team)
	return res.Error
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
	Teams     []Team     `json:"teams" gorm:"ForeignKey:DivisionName;AssociationForeignKey:Name"`
}

func CreateDivision(d Division, DB *gorm.DB) error {
	res := DB.Create(&d)
	return res.Error
}

func FetchDivision(id uint, DB *gorm.DB) (Division, error) {
	var d Division
	res := DB.Preload("Teams").Where("ID = ? AND deleted_at IS NULL", id).First(&d)
	return d, res.Error
}

func FetchDivisions(DB *gorm.DB) ([]Division, error) {
	d := make([]Division, 0)
	res := DB.Preload("Teams").Where("deleted_at IS NULL").Find(&d)
	return d, res.Error
}

func UpdateDivision(d Division, DB *gorm.DB) error {
	res := DB.Where("ID = ? AND deleted_at IS NULL", d.ID).Save(&d)
	return res.Error
}

func DeleteDivision(id int, DB *gorm.DB) error {
	var div Division
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&div)
	return res.Error
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
	res := DB.Create(&p)
	return res.Error
}

func FetchPlayer(id uint, DB *gorm.DB) (Player, error) {
	var p Player
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).First(&p)
	return p, res.Error
}

func FetchPlayers(DB *gorm.DB) ([]Player, error) {
	p := make([]Player, 0)
	res := DB.Where("deleted_at IS NULL").Find(&p)
	return p, res.Error
}

func UpdatePlayer(p Player, DB *gorm.DB) error {
	res := DB.Where("ID = ? AND deleted_at IS NULL", p.ID).Save(&p)
	return res.Error
}

func DeletePlayer(id int, DB *gorm.DB) error {
	var p Player
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&p)
	return res.Error
}
