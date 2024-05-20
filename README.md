# Ticket Booking CLI

This command-line interface (CLI) application allows users to book tickets for an event by providing their first name, last name, email address, and the number of tickets they wish to book. The application provides real-time feedback on the remaining tickets available and maintains a list of users who have successfully booked tickets.

## Features
- **User Input**: Users are prompted to enter their first name, last name, email address, and the number of tickets they wish to book.
- **Input Validation**: The application validates user input to ensure the correctness of entered data. It checks for valid email format and ensures that the number of tickets requested is within the available range.
- **Remaining Tickets**: After each successful booking, the application displays the number of remaining tickets available for the event.
- **User List**: The application maintains a list of users who have successfully booked tickets. It displays this list along with the first names of the users.
- **Ticket Generation**: Simulates ticket generation using goroutines to handle the process asynchronously.

## How to Use
1. Clone the repository to your local machine.
2. Navigate to the directory containing the `main.go` file.
3. Run the command `go run .` to start the application.
4. Follow the on-screen prompts to book tickets for the event.

## Example
- Welcome to Tomorrow land booking application
- 50 tickets remaining out of 50
- Please enter your first name: John
- Please enter your last name: Melony
- Enter your email: Melony@gmail
- Enter the number of tickets: 12
- Here is your detals printed on your ticket [{John Melony Melony@gmail 12}]
- Thank you John Melony for booking 12 tickets, you will receive the tickets on your email Melony@gmail
- 38 remaining out of 50
- Travelers list: [John]
---------------------
- Sending ticket
 - Hello John here is your ticket, John, Melony, Melony@gmail, 12 to Melony@gmail
---------------------
