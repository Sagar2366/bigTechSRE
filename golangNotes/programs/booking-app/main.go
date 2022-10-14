package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference" // conferenceName := "Go Conference" this syntax will only work with variables
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attent the conference.")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to buy : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
