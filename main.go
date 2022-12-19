package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	var allTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	greetUser(conferenceName, conferenceTickets, allTickets)
	
	for {
		// Validation:
		firstName, lastName, email, userTickets := userData(allTickets)
		// Calculation:
		allTickets, bookings = calculationData(allTickets, userTickets, bookings, firstName, lastName)
		// Print output:
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The array type: %T\n", bookings)
		fmt.Printf("The slice length: %v\n", len(bookings))
		fmt.Printf("The first item of slice: %v\n", bookings[0])
		fmt.Printf("User %v %v booked %v tickets, you will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("The first names of bookings are: %v\n", getFirstName(bookings))
		fmt.Printf("%v tickets remaining for %v\n", allTickets, conferenceTickets)
		// The end:
		if allTickets == 0 {
			fmt.Println("Tickets sold out")
			break
		} 
	}
}

func greetUser(name string, tickets int, alltickets uint) {
	fmt.Printf("Hello and welcome to %v booking application!\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", tickets, alltickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName(input []string) []string {
	firstNames := []string{}
	for _, elem := range input {
		names := strings.Fields(elem)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validFirstName(input string) string {
	fmt.Scan(&input)
	for len(input) <= 2 {
		fmt.Printf("The first name %v cannot consist of one or two letters, please enter a valid name.\n", input)
		fmt.Scan(&input)
	}
	return input
}

func validSecondName(input string) string {
	fmt.Scan(&input)
	for len(input) <= 2 {
		fmt.Printf("The last name %v cannot consist of one or two letters, please enter a valid name.\n", input)
		fmt.Scan(&input)
	}
	return input
}

func validEmail(input string) string {
	fmt.Scan(&input)
	for !strings.Contains(input, "@") {
		fmt.Println("Email must contain @")
		fmt.Scan(&input)
	}
	return input
}

func validNumberOfTickets(user uint, all uint) uint {
	fmt.Scan(&user)
	for user > 0 && user > all{
	   fmt.Printf("We only have %v tickets.\nEnter correct number of tickets, please.\n", all )
	   fmt.Scan(&user)
   	}
   	return user	
}

func userData (allTickets uint) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Inputs and validation:
	fmt.Println("Enter your first name, please:")
	firstName = validFirstName(firstName)
	fmt.Println("Enter your last name, please:")
	lastName = validSecondName(lastName)
	fmt.Println("Enter your e-mail adress, please")
	email = validEmail(email)
	fmt.Println("Enter number of tickets")
	userTickets = validNumberOfTickets(userTickets, allTickets)

	return firstName, lastName, email, userTickets
}

func calculationData(allTickets uint, userTickets uint, bookings []string, firstName string, lastName string )(uint, []string) {
	allTickets = allTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	return allTickets, bookings
}