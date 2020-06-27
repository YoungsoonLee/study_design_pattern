package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.payDone {
		fmt.Println("patient already paid")
		c.next.execute(p)
		return
	}
	fmt.Println("patient paying")
	p.payDone = true
	c.next.execute(p)
}

func (c *cashier) setNext(next department) {
	c.next = next
}
