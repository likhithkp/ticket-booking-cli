package main

import (
	"fmt"
	"go-ticket-booking-app/helper"
	"sync"
	"time"
)

const placeName = "Tomorrow land"
const tickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(tickets, firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		firstNames := printFirstNames()
		fmt.Printf("Travelers list: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets sold out, please come back next year")
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
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", placeName)
	fmt.Printf("%v tickets remaining out of %v\n", remainingTickets, tickets)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Here is your detals printed on your ticket %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive the tickets on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining out of %v\n", remainingTickets, tickets)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("Hello %v here is your ticket, %v, %v, %v, %v\n", firstName, firstName, lastName, email, userTickets)
	fmt.Println("---------------------")
	fmt.Printf("Sending ticket\n %v to %v\n", ticket, email)
	fmt.Println("---------------------")
	wg.Done()
}
