package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("conferenceName is %T and conferenceTikets is %T and remainingTickets is %T. \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("welocme to out", conferenceName, "conference application")
	fmt.Printf("we have total of %v tikets and %v.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets from here");

	var userName string
	var userTickets int

	userName = "tharindu"
	userTickets = 10

	fmt.Printf("user name %v and tickets count %v.\n", userName, userTickets)


}  