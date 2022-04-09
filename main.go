package main

import (
	"fmt"
	"strings"
)

// variables shared to main and all functions are package level variables.
//Define variables as LOCAL as posssible

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		bookTicket(userTickets, firstName, lastName, email)

		if isValidName && isValidEmail && isValidTicketNumber {

			firstNames := getFirstNames()
			fmt.Printf("The first name(s) of our booked guest(s) %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("SOLD OUT, MAYBE NEXT TIME!")
				break
			}

		} else {

			if !isValidName {
				fmt.Printf("The first or last name you entered is too short. Please Try Again. \n")
			}
			if !isValidEmail {
				fmt.Printf("Email address does not contain an @ sign. Please Try Again. \n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you entered is invalid. Please Try Again. \n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
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

// no parameters because user input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
