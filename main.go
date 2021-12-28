package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"
const totalTickets = 50
var remainingTickets uint = 50
var firstNames = []string{}

func main() {
	greetUsers()
	for {
		userTickets, firstName, lastName, emailAddress := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)
		if !isValidName {
			fmt.Println("Firstname or Last Name should have more than 2 characters")
		} else if !isValidEmail {
			fmt.Println("Invalid Email Address")
		} else if !isValidTicketNumber {
			fmt.Printf("required tickets %v cannot be more than remaining %v tickets\n", userTickets, remainingTickets)
		} else {
			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["emailAddress"] = emailAddress
			userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
			bookTickets(userData)
			if remainingTickets == 0 {
				fmt.Println("Conference is sold out")
				break
			}
		}
	}
}

func bookTickets(userData map[string]string) {
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", 
				userData["firstName"], userData["lastName"], userData["userTickets"], userData["emailAddress"])
	var userTicks, _ = strconv.ParseUint(userData["userTickets"], 10, 64)
	remainingTickets = remainingTickets - uint(userTicks)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	var bookings []map[string]string	
	bookings = append(bookings, userData)
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

func printFirstNames(bookings []map[string]string) {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v tickets and are %v still available\n",totalTickets, remainingTickets)
}