package mock

import "github.com/thomaspaulin/snc-server-go/snc"

//--------------------------------------------------------------------------------------------------------------------//
// Divisions
//--------------------------------------------------------------------------------------------------------------------//
// DivisionService represents a mock implementation of snc.DivisionService
type DivisionService struct {
	CreateFn			func(d *snc.Division) error
	CreateInvoked		bool

	DivisionFn			func(id int) (*snc.Division, error)
	DivisionInvoked		bool

	DivisionsFn			func() ([]*snc.Division, error)
	DivisionsInvoked	bool

	UpdateFn			func(d *snc.Division) error
	UpdateInvoked		bool

	DeleteFn			func(id int) error
	DeleteInvoked		bool
}

func (ds *DivisionService) CreateDivision(d *snc.Division) error {
	ds.CreateInvoked = true
	return ds.CreateFn(d)
}

func (ds *DivisionService) Division(id int) (*snc.Division, error) {
	ds.DivisionInvoked = true
	return ds.DivisionFn(id)
}

func (ds *DivisionService) Divisions() ([]*snc.Division, error) {
	ds.DivisionsInvoked = true
	return ds.DivisionsFn()
}

func (ds *DivisionService) UpdateDivision(d *snc.Division) error {
	ds.UpdateInvoked = true
	return ds.UpdateFn(d)
}

func (ds *DivisionService) DeleteDivision(id int) error {
	ds.DeleteInvoked = true
	return ds.DeleteFn(id)
}

//--------------------------------------------------------------------------------------------------------------------//
// Teams
//--------------------------------------------------------------------------------------------------------------------//
// TeamService represents a mock implementation of snc.TeamService
type TeamService struct {
	CreateFn			func(t *snc.Team) error
	CreateInvoked		bool

	TeamFn				func(id int) (*snc.Team, error)
	TeamInvoked			bool

	TeamsFn				func() ([]*snc.Team, error)
	TeamsInvoked		bool

	UpdateFn			func(t *snc.Team) error
	UpdateInvoked		bool

	DeleteFn			func(id int) error
	DeleteInvoked		bool
}

func (ts *TeamService) CreateTeam(t *snc.Team) error {
	ts.CreateInvoked = true
	return ts.CreateFn(t)
}

func (ts *TeamService) Team(id int) (*snc.Team, error) {
	ts.TeamInvoked = true
	return ts.TeamFn(id)
}

func (ts *TeamService) Teams() ([]*snc.Team, error) {
	ts.TeamsInvoked = true
	return ts.TeamsFn()
}

func (ts *TeamService) UpdateTeam(t *snc.Team) error {
	ts.UpdateInvoked = true
	return ts.UpdateFn(t)
}

func (ts *TeamService) DeleteTeam(id int) error {
	ts.DeleteInvoked = true
	return ts.DeleteFn(id)
}

//--------------------------------------------------------------------------------------------------------------------//
// Rinks
//--------------------------------------------------------------------------------------------------------------------//
// RinkService represents a mock implementation of snc.RinkService
type RinkService struct {
	CreateFn			func(t *snc.Rink) error
	CreateInvoked		bool

	RinkFn				func(id int) (*snc.Rink, error)
	TeamInvoked			bool

	RinksFn				func() ([]*snc.Rink, error)
	TeamsInvoked		bool

	UpdateFn			func(t *snc.Rink) error
	UpdateInvoked		bool

	DeleteFn			func(id int) error
	DeleteInvoked		bool
}

func (rs *RinkService) CreateRink(t *snc.Rink) error {
	rs.CreateInvoked = true
	return rs.CreateFn(t)
}

func (rs *RinkService) Rink(id int) (*snc.Rink, error) {
	rs.TeamInvoked = true
	return rs.RinkFn(id)
}

func (rs *RinkService) Rinks() ([]*snc.Rink, error) {
	rs.TeamsInvoked = true
	return rs.RinksFn()
}

func (rs *RinkService) UpdateRink(t *snc.Rink) error {
	rs.UpdateInvoked = true
	return rs.UpdateFn(t)
}

func (rs *RinkService) DeleteRink(id int) error {
	rs.DeleteInvoked = true
	return rs.DeleteFn(id)
}