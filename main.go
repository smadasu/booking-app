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
		var userTickets uint
		fmt.Println("Please enter the number of tickets needed")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {
			var firstName string
			var lastName string
			var emailAddress string
			fmt.Println("Please enter your first name")
			fmt.Scan(&firstName)
			fmt.Println("Please enter your last name")
			fmt.Scan(&lastName)
			fmt.Println("Please enter your email address")
			fmt.Scan(&emailAddress)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			var bookings [] string
			bookings = append(bookings, firstName + " " + lastName)
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Fields(booking)[0])
			}
			fmt.Printf("The first names of bookings are = %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Conference is sold out")
				break
			}
		} else {
			fmt.Printf("We only have %v tckets remaining. Please enter a value less than or equal to that\n", remainingTickets)
			continue
		}
	}
}