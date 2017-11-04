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


//--------------------------------------------------------------------------------------------------------------------//
// Rinks
//--------------------------------------------------------------------------------------------------------------------//