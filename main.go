package main

import (
	"fmt"
	"strings"

	"github.com/Akash2023das/ticketbookingapp/helper"
)

const conferenceTickets = 50 //we have total 50 tickets for the conference
var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{} //empty slice

func main() {

	greetUsers() //greeting the users

	for {
		//calling userinput function and storing the return values
		firstName, lastName, email, userTickets := getUserInput()

		//checking the values are correct or not
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//check if the user booked more tickes than remaining
		if isValidName && isValidEmail && isValidTicketNumber {

			//calling function to
			bookTicket(userTickets, firstName, lastName, email)

			//calling function for first name
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are:%v\n", firstNames)

			//if all tickets are booked then break the loop
			if remainingTickets == 0 {
				//end
				fmt.Println("Our conference is booked.Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contains '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

// Greetings the User
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// it will stores the first names of the bookings
func getFirstNames() []string {
	//for storing the first name only
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// asking the user for input
func getUserInput() (string, string, string, uint) {
	//asking user for input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// booking the tickets
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

}
