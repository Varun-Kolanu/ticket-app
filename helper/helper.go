package helper

import "strings"

func Validate(fName, sName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(sName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isUserTicketsValid
}
