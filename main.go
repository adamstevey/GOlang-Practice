package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("There are %v tickets in total.\n", conferenceTickets)
	fmt.Printf("There are %v tickets remaining.\n\n", remainingTickets)

	var bookings []string

	var userName string
	var lastName string
	var userTickets uint

	for {
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("How many ticket would you like?")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, userName+" "+lastName)

			fmt.Printf("Thank you %v %v for buying %v tickets.\n", userName, lastName, userTickets)
			fmt.Printf("There are %v tickets remaing for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Here is a list of bookings:\n %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				break
			}
		} else {
			fmt.Println("Invalid number of tickets, try again.")
		}
	}

}
