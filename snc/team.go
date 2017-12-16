package snc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//--------------------------------------------------------------------------------------------------------------------//
// Team
//--------------------------------------------------------------------------------------------------------------------//
type Team struct {
	ID           uint       `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
	Name         string     `json:"name" gorm:"not null;unique_index;primary_key"`
	LeagueName   string     `json:"leagueName"`
	DivisionName string     `json:"divisionName"`
}

func CreateTeam(t Team, DB *gorm.DB) error {
	team, _ := FetchTeamNamed(t.Name, DB)
	if team.ID != 0 {
		t.ID = team.ID
		return UpdateTeam(t, DB)
	}
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

func FetchTeamNamed(name string, DB *gorm.DB) (Team, error) {
	var team Team
	res := DB.Where("name = ? AND deleted_at IS NULL", name).First(&team)
	return team, res.Error
}

func UpdateTeam(t Team, DB *gorm.DB) error {
	res := DB.Where("id = ? AND deleted_at IS NULL", t.ID).Save(&t)
	return res.Error
}

func DeleteTeam(id uint, DB *gorm.DB) error {
	var team Team
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&team)
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
	Name      string     `json:"name" gorm:"not null;unique_index;primary_key"`
	LeagueName string    `json:"league gorm:"not null"`	  
	Teams     []Team     `json:"teams" gorm:"ForeignKey:DivisionName;AssociationForeignKey:Name"`
}

func CreateDivision(d Division, DB *gorm.DB) error {
	division, _ := FetchDivisionNamed(d.Name, DB)
	if division.ID != 0 {
		d.ID = division.ID
		return UpdateDivision(d, DB)
	}
	res := DB.Create(&d)
	return res.Error
}

func FetchDivision(id uint, DB *gorm.DB) (Division, error) {
	var d Division
	res := DB.Preload("Teams").Where("ID = ? AND deleted_at IS NULL", id).First(&d)
	return d, res.Error
}

func FetchDivisionNamed(name string, DB *gorm.DB) (Division, error) {
	var div Division
	res := DB.Preload("Teams").Where("name = ? AND deleted_at IS NULL", name).First(&div)
	return div, res.Error
}

func FetchDivisions(DB *gorm.DB) ([]Division, error) {
	d := make([]Division, 0)
	res := DB.Preload("Teams").Where("deleted_at IS NULL").Find(&d)
	return d, res.Error
}

func UpdateDivision(d Division, DB *gorm.DB) error {
	res := DB.Where("id = ? AND deleted_at IS NULL", d.ID).Save(&d)
	return res.Error
}

func DeleteDivision(id int, DB *gorm.DB) error {
	var div Division
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&div)
	return res.Error
}

//--------------------------------------------------------------------------------------------------------------------//
// League
//--------------------------------------------------------------------------------------------------------------------//
type League struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name" gorm:"not null;unique_index;primary_key"`
	Divisions     []Division     `json:"divisions" gorm:"ForeignKey:LeagueName;AssociationForeignKey:Name"`
}

func CreateLeague(l League, DB *gorm.DB) error {
	league, _ := FetchLeagueNamed(l.Name, DB)
	if league.ID != 0 {
		l.ID = league.ID
		return UpdateLeague(l, DB)
	}
	res := DB.Create(&l)
	return res.Error
}

func FetchLeague(id uint, DB *gorm.DB) (League, error) {
	var l League
	res := DB.Preload("Leagues").Where("ID = ? AND deleted_at IS NULL", id).First(&l)
	return l, res.Error
}

func FetchLeagueNamed(name string, DB *gorm.DB) (League, error) {
	var l League
	res := DB.Preload("Leagues").Where("name = ? AND deleted_at IS NULL", name).First(&l)
	return l, res.Error
}

func FetchLeagues(DB *gorm.DB) ([]League, error) {
	l := make([]League, 0)
	res := DB.Preload("League").Where("deleted_at IS NULL").Find(&l)
	return l, res.Error
}

func UpdateLeague(l League, DB *gorm.DB) error {
	res := DB.Where("id = ? AND deleted_at IS NULL", l.ID).Save(&l)
	return res.Error
}

func DeleteLeague(id int, DB *gorm.DB) error {
	var l League
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&l)
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
