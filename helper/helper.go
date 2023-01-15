package helper

import "strings"

// it will checks the data is valid or not
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	//checks the user name atleat two characters long or not
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	//check the user email that contains @ or not
	isValidEmail := strings.Contains(email, "@")

	//check for valid ticket
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
