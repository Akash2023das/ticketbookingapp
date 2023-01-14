package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50 //we have total 50 tickets for the conference
	var remainingTickets uint = 50
	bookings := []string{} //empty slice

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
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

		//checks the user name atleat two characters long or not
		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		//check the user email that contains @ or not
		isValidEmail := strings.Contains(email, "@")

		//check for valid ticket
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//check if the user booked more tickes than remaining
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			//for storing the first name only
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are:%v\n", firstName)

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
func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
