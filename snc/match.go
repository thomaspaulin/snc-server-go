package snc

import (
	"fmt"
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

type Pagination struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type DateRange struct {
	Start time.Time `form:"from" time_format:"2006-01-02"`
	End   time.Time `form:"to" time_format:"2006-01-02"`
}

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
	var awayID, homeID, rinkID uint
	away, _ := FetchTeamNamed(m.Away.Name, DB)
	awayID = away.ID
	home, _ := FetchTeamNamed(m.Home.Name, DB)
	homeID = home.ID
	rink, _ := FetchRinkNamed(m.Rink.Name, DB)
	rinkID = rink.ID

	// For some reason IDs aren't being fetched so put this hack in place
	var res *gorm.DB
	if awayID != 0 && homeID != 0 && rinkID != 0 {
		var match Match
		res = DB.Where("start = ? AND away_id = ? AND home_id = ?", m.Start, awayID, homeID).
			Find(&match)
		if match.ID != 0 {
			m.ID = match.ID
			if m.AwayID == 0 {
				m.AwayID = awayID
				m.Away.ID = awayID
			}
			if m.HomeID == 0 {
				m.HomeID = homeID
				m.Home.ID = homeID
			}
			if m.RinkID == 0 {
				m.RinkID = rinkID
				m.Rink.ID = rinkID
			}
			// found match in DB
			res = DB.Save(&m)
		} else {
			// didn't find match but did find the teams
			m.AwayID = awayID
			m.Away.ID = awayID
			m.HomeID = homeID
			m.Home.ID = homeID
			m.RinkID = rinkID
			m.Rink.ID = rinkID
			res = DB.Create(&m)
		}
	} else {
		if m.AwayID == 0 {
			m.AwayID = awayID
			m.Away.ID = awayID
		}
		if m.HomeID == 0 {
			m.HomeID = homeID
			m.Home.ID = homeID
		}
		if m.RinkID == 0 {
			m.RinkID = rinkID
			m.Rink.ID = rinkID
		}
		res = DB.Create(&m)
	}
	return res.Error
}

func FetchMatch(id uint, DB *gorm.DB) (Match, error) {
	m := Match{}
	res := DB.Preload("Away").
		Preload("Home").
		Preload("Rink").
		Where("id = ? AND deleted_at IS NULL", id).First(&m)
	return m, res.Error
}

func FetchMatches(dates DateRange, pagination Pagination, DB *gorm.DB) ([]Match, error) {
	if pagination.Offset == 0 || pagination.Offset < -1 {
		fmt.Println("Pagination offset not set or invalid. Cancelling.")
		// if offset not set or invalid cancel the offset (-1 in gorm)
		pagination.Offset = -1
	}
	if pagination.Limit == 0 || pagination.Limit < -1 {
		fmt.Println("Pagination limit not set or invalid. Cancelling.")
		// if limit not set or invalid cancel the limit (-1 in gorm)
		pagination.Limit = -1
	}

	query := "deleted_at IS NULL"
	dateRange := make([]interface{}, 0)
	if !dates.Start.IsZero() {
		query += " AND start >= ?"
		dateRange = append(dateRange, dates.Start)
	}
	if !dates.End.IsZero() {
		query += " AND start < ?"
		dateRange = append(dateRange, dates.End)
	}

	m := make([]Match, 0)
	res := DB.Preload("Away").
		Preload("Home").
		Preload("Rink").
		Where(query, dateRange...).
		Order("start asc").
		Offset(pagination.Offset).
		Limit(pagination.Limit).
		Find(&m)
	return m, res.Error
}

func UpdateMatch(m Match, DB *gorm.DB) error {
	res := DB.Where("deleted_at IS NULL").Save(&m)
	return res.Error
}

func DeleteMatch(id uint, DB *gorm.DB) error {
	var m Match
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&m)
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
		Where("id = ? AND deleted_at IS NULL", id).First(&g)
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
	res := DB.Where("id = ? AND deleted_at IS NULL", id).Delete(&g)
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
