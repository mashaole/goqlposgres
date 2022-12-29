package domain

import "goqlposgress/postgres"

type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, MeetupsRepo: meetupsRepo}
}
