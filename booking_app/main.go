package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
    // var bookings = [50]string  --> this is an array
	// bookings[1] = "Nicole"
	// bookings[0] = firstName + " " + lastName
	bookings := []string{}
	// this is a slice


	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

    for {
	    var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fnt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking:= range bookings {
			var names = string.Fields(booking)
			firstNames = append(firtNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}