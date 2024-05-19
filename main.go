package main

import (
	"fmt"
	"strings"
)

const placeName = "Tomorrow land"
const tickets uint = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(tickets, firstName, lastName, email, userTickets)

			firstNames := printFirstNames()
			fmt.Printf("Travelers list: %v\n", firstNames)

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

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", placeName)
	fmt.Printf("%v tickets remaining out of %v\n", remainingTickets, tickets)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(tickets uint, firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive the tickets on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining out of %v\n", remainingTickets, tickets)
}
