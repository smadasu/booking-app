package main
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceName is %T, totalTickets is %T and remainingTickets is %T\n", conferenceName, totalTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and are %v still available\n",totalTickets, remainingTickets, )
	fmt.Println("Get your tickets here to attend")
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
	fmt.Printf("Thank you %v %v for booking  %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, emailAddress, userTickets)
	remainingTickets = totalTickets - userTickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}