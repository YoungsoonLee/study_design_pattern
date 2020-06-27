package main

type userCollection struct {
	users []*user
}

func (u *userCollection) createIteration() iterator {
	return &userIterator{
		users: u.users,
	}
}
