package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference" // conferenceName := "Go Conference" this syntax will only work with variables
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string
	var bookingsSlice []string

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attent the conference.")

	for {
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
		bookings[0] = firstName + " " + lastName
		bookingsSlice = append(bookingsSlice, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("Bookings array is : %v\n", bookings)
		fmt.Printf("First element in Bookings array is : %v\n", bookings[0])
		fmt.Printf("Type of a Bookings array is : %T\n", bookings)
		fmt.Printf("Lenght of a Bookings array is : %v\n", len(bookings))

		fmt.Printf("Bookings slice is : %v\n", bookingsSlice)
		fmt.Printf("First element in Bookings slice is : %v\n", bookingsSlice[0])
		fmt.Printf("Type of a Bookings slice is : %T\n", bookingsSlice)
		fmt.Printf("Lenght of a Bookings slice is : %v\n", len(bookingsSlice))

		keepFirstNames := []string{}
		for _, booking := range bookingsSlice {
			name := strings.Fields(booking)
			keepFirstNames = append(keepFirstNames, name[0])

		}
		fmt.Printf("First names of the bookings are: %v\n", keepFirstNames)
	}

}
