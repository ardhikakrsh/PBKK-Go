package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// fmt.Println("Welcome to", conferenceName, "booking application!")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get your tickets now!")

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets now!\n")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v. Your tickets are confirmed. Please check your email %v for the confirmation.\n", firstName, lastName, userTickets, conferenceName, email)

}