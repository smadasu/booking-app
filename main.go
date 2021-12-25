package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	firstNames := []string{}
	for {
		fmt.Printf("We have total of %v tickets and are %v still available\n",totalTickets, remainingTickets)
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
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		var bookings [] string
		bookings = append(bookings, firstName + " " + lastName)
		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Fields(booking)[0])
		}
		fmt.Printf("The first names of bookings are = %v\n", firstNames)
	}
}