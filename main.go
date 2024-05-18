package main

import (
	"fmt"
	"strings"
)

func main() {
	const placeName = "Tomorrow land"
	const tickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to the %v ticket booking system\n", placeName)
	fmt.Printf("%v tickets remaining out of %v\n", remainingTickets, tickets)
	fmt.Printf("Book your tickets below to travel to %v\n", placeName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive the tickets on your email %v\n", firstName, lastName, userTickets, email)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Travelers list: %v\n", firstNames)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, placeName)
	}
}
