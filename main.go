package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("welocme to out", conferenceName, "conference application")
	fmt.Printf("we have total of %v tikets and %v.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets from here");

	var userName string
	var lastName string
	var email string
	var userTickets int

	// ask for the user input
	fmt.Print("enter your first name: ")
	fmt.Scan(&userName)

	fmt.Print("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("enter your email: ")
	fmt.Scan(&email)

	fmt.Print("how many tickets do u want : ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank your for %v %v for book %v tickets and receipt will be deliver %v address\n", userName, lastName, userTickets, email)


}  