package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}

	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}
	iterator := userCollection.createIteration()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User us %+v\n", user)
	}
}
