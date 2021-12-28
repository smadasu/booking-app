package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"
const totalTickets = 50
var remainingTickets uint = 50
var firstNames = []string{}

func main() {
	greetUsers()
	for {
		userTickets, firstName, lastName, emailAddress := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets)
		if !isValidName {
			fmt.Println("Firstname or Last Name should have more than 2 characters")
		} else if !isValidEmail {
			fmt.Println("Invalid Email Address")
		} else if !isValidTicketNumber {
			fmt.Printf("required tickets %v cannot be more than remaining %v tickets", userTickets, remainingTickets)
		} else {
			bookTickets(firstName, lastName, userTickets, emailAddress)
			if remainingTickets == 0 {
				fmt.Println("Conference is sold out")
				break
			}
		}
	}
}

func bookTickets(firstName string, lastName string, userTickets uint, emailAddress string) {
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	var bookings []string
	bookings = append(bookings, firstName+" "+lastName)
	printFirstNames(bookings)
	fmt.Printf("The first names of bookings are = %v\n", firstNames)
}

func getUserInput() (uint, string, string, string) {
	var userTickets uint
	fmt.Println("Please enter the number of tickets needed")
	fmt.Scan(&userTickets)
	var firstName string
	var lastName string
	var emailAddress string
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address")
	fmt.Scan(&emailAddress)
	return userTickets, firstName, lastName, emailAddress
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func printFirstNames(bookings []string) {
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v tickets and are %v still available\n",totalTickets, remainingTickets)
}