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

    isValidName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email, "@")
    isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

    if isValidName && isValidEmail && isValidTicketNumber {
      remainingTickets = remainingTickets - userTickets
		  bookings = append(bookings, firstName+" "+lastName)

		  fmt.Printf("Thank you %v %v for booking %v tickets, you will receive the tickets on your email %v\n", firstName, lastName, userTickets,email)

		  firstNames := []string{}
	  	for _, booking := range bookings {
			  var names = strings.Fields(booking)
			  firstNames = append(firstNames, names[0])
		  }

		  fmt.Printf("Travelers list: %v\n", firstNames)
		  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, placeName)

      if remainingTickets == 0 {
        fmt.Println("Tickets sold out, please come back next year")
        break
      }
    } else {
        if !isValidName {
          fmt.Println("Invalid name, first name and last name must be more than 2 characters")
        }
        if !isValidEmail {
          fmt.Println("Invalid email, please ensure that your email has @ in it")
        }
        if !isValidTicketNumber {
          fmt.Printf("Invalid ticket counts, you can book between 1 to %v\n", remainingTickets)
        }
    }
  }
}
