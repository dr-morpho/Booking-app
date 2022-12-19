package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	conferenceTickets := 50
	var allTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, allTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name, please:")
		fmt.Scan(&firstName)
		for len(firstName) <= 2 {
			fmt.Printf("The first name %v cannot consist of one or two letters, please enter a valid name.\n", firstName)
			fmt.Scan(&firstName)
		}

		fmt.Println("Enter your last name, please:")
		fmt.Scan(&lastName)
		for len(lastName) <= 2 {
			fmt.Printf("The last name %v cannot consist of one or two letters, please enter a valid name.\n", lastName)
			fmt.Scan(&lastName)
		}

		fmt.Println("Enter your e-mail adress, please")
		fmt.Scan(&email)
		for !strings.Contains(email, "@") {
			fmt.Println("Email must contain @")
			fmt.Scan(&email)
		}

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
 		for userTickets > 0 && userTickets > allTickets {
			fmt.Printf("We only have %v tickets.\nEnter correct number of tickets, please.\n", allTickets )
			fmt.Scan(&userTickets)
		}		

		allTickets = allTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
		
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The array type: %T\n", bookings)
		fmt.Printf("The slice length: %v\n", len(bookings))
		fmt.Printf("The first item of slice: %v\n", bookings[0])
		fmt.Printf("User %v %v booked %v tickets, you will receive a confirmation email at %v).\n", firstName, lastName, userTickets, email)
	
		firstNames := []string{}
		for _, elem := range bookings {
			names := strings.Fields(elem)
			firstNames = append(firstNames, names[0])
		}
			
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		fmt.Printf("%v tickets remaining for %v\n", allTickets, conferenceTickets)

		if allTickets == 0 {
			fmt.Println("Tickets sold out")
			break
		} 
	}
}

	