package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "GopherCon"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	userEmail       string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, userEmail, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(remainingTickets, firstName, lastName, userEmail, userTickets)

	if isValidEmail && isValidName && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, userEmail)

		wg.Add(1) // adds a goroutine thread to wait for the function to execute before exiting main goroutine thread
		go sendTicket(userTickets, firstName, lastName, userEmail)
		// THE GO KEYWORD MAKES IT CONCURRENT FUNCTION

		firstNames := getFirstNames()
		fmt.Printf("All the bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("%v is booked out! Come back next year.\n", conferenceName)
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name or Last Name is too short")
		}
		if !isValidEmail {
			fmt.Println("Input email does not have @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid number of tickets")
		}
	}
	wg.Wait() // blocks the main function quitting until WaitGroup counter is 0
}

func greetUsers() {
	fmt.Printf("Welcome to our %v tickets booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var userTickets uint
	var firstName string
	var lastName string
	var userEmail string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets -= userTickets

	// create a struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		userEmail:       userEmail,
		numberOfTickets: userTickets,
	}

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["userEmail"] = userEmail
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// we can only have one datatype to a map

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	fmt.Println("Processing your ticket(s)....")
	time.Sleep(10 * time.Second) // To simulate the task of generating a ticket and sending email

	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("<<---------------->>")
	fmt.Printf("Sending ticket:\n %v\n to email address %v\n", ticket, userEmail)
	fmt.Println("<<---------------->>")
	wg.Done()
}
