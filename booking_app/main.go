package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")

	// bookings[1] = "Nicole"

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// var bookings = [50]string  --> this is an array
	var bookings = []string
	// this is a slice
	

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter how many tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
    // bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fnt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

    fmt.Printf("These are all our bookings: %v\n", bookings)
}
