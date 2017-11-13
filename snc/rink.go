package snc

const (
	Avondale = "Avondale"
	Botany = "Botany"
)

type Rink struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
}

type RinkService interface {
	CreateRink(r *Rink) error
	Rink(id int) (*Rink, error)
	Rinks() ([]*Rink, error)
	UpdateRink(r *Rink) error
	DeleteRink(id int) error
}
