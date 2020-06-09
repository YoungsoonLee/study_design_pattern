package main

type connecntion struct {
	id string
}

func (c *connecntion) getID() string {
	return c.id
}
