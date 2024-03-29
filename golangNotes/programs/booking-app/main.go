package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var conferenceName string = "Go Conference" // conferenceName := "Go Conference" this syntax will only work with variables
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string
var bookingsMap = make([]map[string]string, 0) // Initialising list of maps, we need to provide intiale size - here we used 0

var bookingsStruct = make([]UserData, 0)

type UserData struct {
	firstName1       string
	lastName1        string
	email1           string
	numberOfTickets1 uint
}

var waitGroup = sync.WaitGroup{}

func main() {
	/*
		Function level variables
		var conferenceName string = "Go Conference" // conferenceName := "Go Conference" this syntax will only work with variables
		const conferenceTickets int = 50
		var remainingTickets uint = 50
		var bookings [50]string
		var bookingsSlice []string
	*/

	greetUsers()

	//for remainingTickets > 0 {

	firstName, lastName, email, userTickets := getUserInput()
	getFirstNames := printFirstNames()
	FirstNamesFromMap := getFirstNamesFromMap()

	fmt.Printf("First names of the bookings are: %v\n", getFirstNames)
	fmt.Printf("First names of the bookings maps are: %v\n", FirstNamesFromMap)

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		waitGroup.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		noRemainingTickets := remainingTickets == 0
		if noRemainingTickets {
			fmt.Printf("Our conference tickets are sold out, please come back next year.")
			//	break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't have @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("The number of tickets you enetred are incorrect.")
		}
		fmt.Println("Your input data is invalid, please try again !!!")
	}
	//	}

	city := "London"
	switch city {
	case "NY":
		fmt.Println("NY")
	case "London":
		fmt.Println("London")
	case "Mumbai", "Pune":
		fmt.Println("Mumbai and Pune")
	default:
		fmt.Println("No valid city selected")
	}

	waitGroup.Wait()
}

func greetUsers() {

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attent the conference.")

}

func printFirstNames() []string {
	getFirstNames := []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking)
		getFirstNames = append(getFirstNames, name[0])

	}
	return getFirstNames
}

func getFirstNamesFromMap() []string {
	firstNames := []string{}

	for _, bookings := range bookingsStruct {
		firstNames = append(firstNames, bookings.firstName1)
	}
	return firstNames
}

func getFirstNamesFromStruct() []string {
	firstNames := []string{}

	for _, bookings := range bookingsMap {
		firstNames = append(firstNames, bookings["firstName"])
	}
	return firstNames
}
func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // Converting uint to string

	userDataStruct := UserData{
		firstName1:       firstName,
		lastName1:        lastName,
		email1:           email,
		numberOfTickets1: userTickets,
	}
	bookings = append(bookings, firstName+" "+lastName)

	bookingsMap = append(bookingsMap, userData)

	bookingsStruct = append(bookingsStruct, userDataStruct)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("Bookings array is : %v\n", bookings)
	fmt.Printf("First element in Bookings array is : %v\n", bookings[0])
	fmt.Printf("Type of a Bookings array is : %T\n", bookings)
	fmt.Printf("Lenght of a Bookings array is : %v\n", len(bookings))

	fmt.Printf("List of bookings is %v\n", bookingsMap)
	fmt.Printf("List of bookings form struct is %v\n", bookingsStruct)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Millisecond)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("####################################################")
	fmt.Printf("Sending %v tickets to email address%v\n", ticket, email)
	fmt.Println("####################################################")

	waitGroup.Done()
}
