package main

import (
	"booking-app/utils"
	"fmt"
)


func main() {
	const conferenceTickets int = 50
	var allTickets uint = 50
	conferenceName := "Go Conference"
	var bookings = make([]map[string]string, 0)

	utils.GreetUser(conferenceName, conferenceTickets, allTickets)
	
	for {
		// Validation:
		firstName, lastName, email, userTickets := utils.UserData(allTickets)
		// Calculation:
		allTickets, bookings = utils.CalculationData(allTickets, userTickets, bookings, firstName, lastName)
		// Print output:
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The array type: %T\n", bookings)
		fmt.Printf("The slice length: %v\n", len(bookings))
		fmt.Printf("The first item of slice: %v\n", bookings[0])
		fmt.Printf("User %v %v booked %v tickets, you will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("The first names of bookings are: %v\n", utils.GetFirstName(bookings))
		fmt.Printf("%v tickets remaining for %v\n", allTickets, conferenceTickets)
		// The end:
		if allTickets == 0 {
			fmt.Println("Tickets sold out")
			break
		} 
	}
}
