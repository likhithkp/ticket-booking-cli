package main

import "fmt"

func main() {
	const placeName = "Tomorrow land"
	const tickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v ticket booking system\n", placeName)
	fmt.Printf("%v tickets remaining out of %v\n", remainingTickets, tickets)
	fmt.Printf("Please book your tickets here to travel to %v\n", placeName)

}
