package main

import (
	"fmt"
	"strconv"
)

// variables shared to main and all functions are package level variables.
//Define variables as LOCAL as posssible

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
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
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

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

	// create map for user
	// define key value string int ext but must have same value
	// convert uint to string using strconv.FormatUint(uint64())

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Turn bookings list from slice of strings to slice of maps
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
