package models

type PersonOrOrganization interface {
	getId() uint
	getName() string
	getEmail() string
}

func (m Thing) getID() uint {
	return m.ID
}
