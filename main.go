package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const totalTickets = 50
var remainingTickets uint = 50
var firstNames = []string{}
var bookings []UserData
var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName string
	emailAddress string
	userTickets uint
}
	
func main() {
	greetUsers()
	userTickets, firstName, lastName, emailAddress := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)
	if !isValidName {
		fmt.Println("Firstname or Last Name should have more than 2 characters")
	} else if !isValidEmail {
		fmt.Println("Invalid Email Address")
	} else if !isValidTicketNumber {
		fmt.Printf("required tickets %v cannot be more than remaining %v tickets\n", userTickets, remainingTickets)
	} else {
		var userData = UserData {
			firstName: firstName,
			lastName: lastName,
			emailAddress: emailAddress,
			userTickets: userTickets,
		}
		bookTickets(userData)
		if remainingTickets == 0 {
			fmt.Println("Conference is sold out")
		}
	}
	wg.Wait()
}

func bookTickets(userData UserData) {
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", 
				userData.firstName, userData.lastName, userData.userTickets, userData.emailAddress)
	wg.Add(1)
	go sendTicket(userData)
	remainingTickets = remainingTickets - userData.userTickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	bookings = append(bookings, userData)
	printFirstNames()
	fmt.Printf("The first names of bookings are = %v\n", firstNames)
	fmt.Printf("The bookings are = %v\n", bookings)
}

func getUserInput() (uint, string, string, string) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address")
	fmt.Scan(&emailAddress)
	fmt.Println("Please enter the number of tickets needed")
	fmt.Scan(&userTickets)
	return userTickets, firstName, lastName, emailAddress
}

func printFirstNames() {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v tickets and are %v still available\n",totalTickets, remainingTickets)
}

func sendTicket(userData UserData) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userData.userTickets, userData.firstName, userData.lastName)
	fmt.Println("*****************")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, userData.emailAddress)
	fmt.Println("*****************")
	wg.Done()
}