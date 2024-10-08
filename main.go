package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("conferenceTickets is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTickets, remainingTickets,conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// var bookings [50]string // This is an array
	var bookings []string // This is a slice

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

	remainingTickets -= userTickets
	// bookings[0] = firstName + " " + lastName // Syntax for array
	bookings = append(bookings, firstName + " " + lastName) // Syntax for slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))
}
