package domain

type Enum interface {
	IsValid() bool
}

type BookCategory int

const (
	Action BookCategory = iota
	Romance
	Mistery
	SciFi
	Contemporary
)

func (b BookCategory) IsValid() bool {
	switch b {
	case Action, Romance, Mistery, SciFi, Contemporary:
		return true
	}
	return false
}
