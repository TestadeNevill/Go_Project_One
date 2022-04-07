package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			printFirstNames(bookings)

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

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first name(s) of our booked guest(s) %v\n", firstNames)
}
