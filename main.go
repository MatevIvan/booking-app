package main

// repo for youtube video: https://gitlab.com/nanuchi/go-full-course-youtube
// youtube video link: https://www.youtube.com/watch?v=yyUHQIec83I

import (
	"fmt"
	"strings"
)

	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets = conferenceTickets
	var bookings []string // This is a slice

func main() {


	greetUsers()

	for remainingTickets > 0 {
		
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()			
			fmt.Printf("The first names of the bookins are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All the tickets have been sold out. Join us next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email you enetered does not conatin \"@\" symbol.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets entered is invalid.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n",conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
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
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of ticket you want to purchase: ")
	fmt.Scan(&userTickets)
	println()

	return firstName, lastName, email, userTickets
}

func bookTickets( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName) 

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
}