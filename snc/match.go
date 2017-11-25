package snc

import (
	"github.com/jinzhu/gorm"
	"time"
)

// todo handle the errors properly
const (
	PracticeGame = "PR"
	RegularGame  = "RS"
	PlayoffGame  = "PO"

	RegularGoal     = "RG"
	PowerPlayGoal   = "PP"
	ShortHandedGoal = "SH"

	Upcoming = "Upcoming"
	Underway = "Underway"
	Over     = "Over"
)

//-----------------------------------------------//
// Match
//-----------------------------------------------//
type Match struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	// Datetime of the match start in UTC
	Start        time.Time `json:"start" gorm:"not null;unique_index:idx_start_away_home"`
	Season       int       `json:"season"`
	Status       string    `json:"status"`
	DivisionName string    `json:"divisionName"`
	Away         Team      `json:"away" `
	Home         Team      `json:"home"`
	AwayID       uint      `json:"-" gorm:"not null;unique_index:idx_start_away_home"`
	HomeID       uint      `json:"-" gorm:"not null;unique_index:idx_start_away_home"`
	AwayScore    uint      `json:"awayScore"`
	HomeScore    uint      `json:"homeScore"`
	Rink         Rink      `json:"rink" gorm:"not null"`
	RinkID       uint      `json:"-"`
	//Goals	   []MatchGoal `json:"goals,omitempty"`
}

func CreateMatch(m Match, DB *gorm.DB) error {
	res := DB.Create(&m)
	return res.Error
}

func FetchMatch(id uint, DB *gorm.DB) (Match, error) {
	m := Match{}
	res := DB.Preload("Away").
		Preload("Home").
		Where("ID = ? AND deleted_at IS NULL", id).First(&m)
	return m, res.Error
}

func FetchMatches(DB *gorm.DB) ([]Match, error) {
	m := make([]Match, 0)
	res := DB.Preload("Away").
		Preload("Away.Division").
		Preload("Home").
		Preload("Home.Division").
		Preload("Rink").Where("deleted_at IS NULL").Find(&m)
	return m, res.Error
}

func UpdateMatch(m Match, DB *gorm.DB) error {
	res := DB.Where("deleted_at IS NULL").Save(&m)
	return res.Error
}

func DeleteMatch(id uint, DB *gorm.DB) error {
	var m Match
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&m)
	return res.Error
}

//-----------------------------------------------//
// Goal
//-----------------------------------------------//
type Goal struct {
	gorm.Model
	GoalType string `json:"goalType"`
	Team     Team   `json:"team"`
	TeamID   uint   `json:"-"`
	Period   uint   `json:"period"`
	// Seconds left in the period when the goal was scored
	Time     uint   `json:"time"`
	Scorer   Player `json:"scoredBy"`
	ScorerID uint   `json:"-"`
	//AssistedBy []Player `json:"assistedBy" gorm:"ForeignKey:ID"`
}

func CreateGoal(g Goal, DB *gorm.DB) error {
	res := DB.Create(&g)
	return res.Error
}

func FetchGoal(id uint, DB *gorm.DB) (Goal, error) {
	g := Goal{}
	res := DB.Preload("Team").
		Preload("Scorer").
		Where("ID = ? AND deleted_at IS NULL", id).First(&g)
	return g, res.Error
}

func FetchGoals(DB *gorm.DB) ([]Goal, error) {
	m := make([]Goal, 0)
	res := DB.Where("deleted_at IS NULL").Find(&m)
	return m, res.Error
}

func UpdateGoal(g Goal, DB *gorm.DB) error {
	res := DB.Where("deleted_at IS NULL").Save(&g)
	return res.Error
}

func DeleteGoal(id uint, DB *gorm.DB) error {
	var g Goal
	res := DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&g)
	return res.Error
}

//type MatchGoal struct {
//  ID        uint `gorm:"primary_key"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//  DeletedAt *time.Time `json:"-" sql:"index"`
//	MatchID uint	`gorm:"index, primary_key"`
//	GoalID uint		`gorm:"index, primary_key"`
//}

////-----------------------------------------------//
//// Penalty
////-----------------------------------------------//
//type Penalty struct {
//	ID     uint   `json:"id"`
//	Team   string `json:"team"`
//	Period uint   `json:"period"`
//	// Seconds left in the period when the penalty was incurred
//	Time uint `json:"time"`
//	// Name of the penalty
//	Offense string `json:"offense"`
//	// ID of the offender
//	Offender string `json:"offender"`
//	// Penalty Infraction Minutes
//	PIM uint `json:"pim"`
//}
