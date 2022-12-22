package utils

import (
	"fmt"
	"strings"
)

type DataUsers struct {
	firstName string
	lastName string 
	userTickets uint
}

func GreetUser(name string, tickets int, alltickets uint) {
	fmt.Printf("Hello and welcome to %v booking application!\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", tickets, alltickets)
	fmt.Println("Get your tickets here to attend")
}

func GetFirstName(input []DataUsers) []string {
	firstNames := []string{}
	for _, elem := range input {
		firstNames = append(firstNames, elem.firstName)
	}
	return firstNames
}

func ValidFirstName(input string) string {
	fmt.Scan(&input)
	for len(input) <= 2 {
		fmt.Printf("The first name %v cannot consist of one or two letters, please enter a valid name.\n", input)
		fmt.Scan(&input)
	}
	return input
}

func ValidSecondName(input string) string {
	fmt.Scan(&input)
	for len(input) <= 2 {
		fmt.Printf("The last name %v cannot consist of one or two letters, please enter a valid name.\n", input)
		fmt.Scan(&input)
	}
	return input
}

func ValidEmail(input string) string {
	fmt.Scan(&input)
	for !strings.Contains(input, "@") {
		fmt.Println("Email must contain @")
		fmt.Scan(&input)
	}
	return input
}

func ValidNumberOfTickets(user uint, all uint) uint {
	fmt.Scan(&user)
	for user > 0 && user > all{
	   fmt.Printf("We only have %v tickets.\nEnter correct number of tickets, please.\n", all )
	   fmt.Scan(&user)
   	}
   	return user	
}

func UserData (allTickets uint) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Inputs and validation:
	fmt.Println("Enter your first name, please:")
	firstName = ValidFirstName(firstName)
	fmt.Println("Enter your last name, please:")
	lastName = ValidSecondName(lastName)
	fmt.Println("Enter your e-mail adress, please")
	email = ValidEmail(email)
	fmt.Println("Enter number of tickets")
	userTickets = ValidNumberOfTickets(userTickets, allTickets)

	return firstName, lastName, email, userTickets
}

func CalculationData(allTickets uint, userTickets uint, bookings []DataUsers, firstName string, lastName string )(uint, []DataUsers) {
	allTickets = allTickets - userTickets
	var userData = DataUsers {
		firstName: firstName,
		lastName: lastName,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	return allTickets, bookings
}