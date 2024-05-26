package main

import (
	"fmt"
	"strconv"
	"ticket-app/helper"
)

const conferenceTickets = 50

var remainingTickets uint = conferenceTickets
var conferenceName = "Go Conference"
var bookings []map[string]string

func main() {

	greetUsers()

	for {

		fName, sName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isUserTicketsValid := helper.Validate(fName, sName, email, userTickets, remainingTickets)

		if !(isValidName && isValidEmail && isUserTicketsValid) {
			if !isValidName {
				fmt.Println("First Name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isUserTicketsValid {
				fmt.Println("Tickets are invalid")
			}
			fmt.Println("Your input data is invalid")
			continue
		}

		bookTickets(fName, sName, email, userTickets)

		remainingTickets -= userTickets

		if noTicketsRemaining := remainingTickets == 0; noTicketsRemaining {
			fmt.Println("Conference is booked out")
			break
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %s conference\n", conferenceName)
	fmt.Printf("We have total of %d tickets and remaining tickets are %d\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func firstNames() []string {
	var fNames []string
	for _, booking := range bookings {
		fNames = append(fNames, booking["fName"])
	}
	return fNames
}

func getUserInput() (fName, sName, email string, userTickets uint) {
	fmt.Print("Enter your first name: ")
	fmt.Scan(&fName)

	fmt.Print("Enter your second name: ")
	fmt.Scan(&sName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return
}

func bookTickets(fName, sName, email string, userTickets uint) {

	var userData = make(map[string]string)
	userData["fName"] = fName
	userData["sName"] = sName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %s for booking %d tickets. You will receive confirmation email at %s\n", fName, userTickets, email)
	fmt.Printf("Remaining tickets are: %d\n", remainingTickets)
	fmt.Printf("Bookings: %v\n", firstNames())
}
