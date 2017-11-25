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
	gorm.Model
	// Datetime of the match start in UTC
	Start      time.Time `json:"start" gorm:"not null;unique_index:idx_start_away_home"`
	Season     int       `json:"season"`
	Status     string    `json:"status"`
	Division   Division  `json:"division"`
	DivisionID uint      `json:"-"`
	Away       Team      `json:"away" gorm:"not null;unique_index:idx_start_away_home"`
	Home       Team      `json:"home" gorm:"not null;unique_index:idx_start_away_home"`
	AwayID     uint      `json:"-"`
	HomeID     uint      `json:"-"`
	AwayScore  uint      `json:"awayScore"`
	HomeScore  uint      `json:"homeScore"`
	Rink       Rink      `json:"rink" gorm:"not null"`
	RinkID     uint      `json:"-"`
	//Goals	   []MatchGoal `json:"goals,omitempty"`
}

func CreateMatch(m Match, DB *gorm.DB) error {
	DB.Create(&m)
	return DB.Error
}

func FetchMatch(id uint, DB *gorm.DB) (Match, error) {
	m := Match{}
	DB.Preload("Division").
		Preload("Away").
		Preload("Home").
		Where("ID = ? AND deleted_at IS NULL", id).First(&m)
	return m, DB.Error
}

func FetchMatches(DB *gorm.DB) ([]Match, error) {
	m := make([]Match, 0)
	DB.Preload("Division").
		Preload("Away").
		Preload("Away.Division").
		Preload("Home").
		Preload("Home.Division").
		Preload("Rink").Where("deleted_at IS NULL").Find(&m)
	return m, DB.Error
}

func UpdateMatch(m Match, DB *gorm.DB) error {
	DB.Where("deleted_at IS NULL").Save(&m)
	return DB.Error
}

func DeleteMatch(id uint, DB *gorm.DB) error {
	var m Match
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&m)
	return DB.Error
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
	DB.Create(&g)
	return DB.Error
}

func FetchGoal(id uint, DB *gorm.DB) (Goal, error) {
	g := Goal{}
	DB.
		Preload("Team").
		Preload("Scorer").
		Where("ID = ? AND deleted_at IS NULL", id).First(&g)
	return g, DB.Error
}

func FetchGoals(DB *gorm.DB) ([]Goal, error) {
	m := make([]Goal, 0)
	DB.Where("deleted_at IS NULL").Find(&m)
	return m, DB.Error
}

func UpdateGoal(g Goal, DB *gorm.DB) error {
	DB.Where("deleted_at IS NULL").Save(&g)
	return DB.Error
}

func DeleteGoal(id uint, DB *gorm.DB) error {
	var g Goal
	DB.Where("ID = ? AND deleted_at IS NULL", id).Delete(&g)
	return DB.Error
}

//type MatchGoal struct {
//	gorm.Model
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
