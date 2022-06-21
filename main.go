package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GoConf"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name.")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name.")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email.")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would like.")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is fully booked. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short. Please try again.")
			}

			if !isValidEmail {
				fmt.Println("Email address must contain @ symbol. Please try again.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid. Please try again.")
			}
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
