package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	greetUsers(conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	var bookings []map[string]string
	for true {
		firstname, lastname, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)
		if !isValidName {
			fmt.Printf("Entered name is too short, try again !\n")
		} else if !isValidEmail {
			fmt.Printf("Entered Email is not in the correct format, try again !\n")
		} else if !isValidTicketNumber {
			fmt.Printf("Invalid ticket-number, try again !\n")
		} else {
			bookTicket(&remainingTickets, userTickets, &bookings, firstname, lastname, email, conferenceName)
			if remainingTickets == 0 {
				fmt.Printf("Our conference has ended, please come next year !\n")
				break
			}
		}
	}
	fmt.Printf("%v", bookings)
}

func bookTicket(remainingTickets *uint, userTickets uint, bookings *[]map[string]string, firstname string, lastname string, email string, conferenceName string) {
	var userData = make(map[string]string)
	userData["firstname"] = firstname
	userData["lastname"] = lastname
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*remainingTickets -= userTickets
	*bookings = append(*bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", *remainingTickets, conferenceName)
}

func greetUsers(conference_name string) {
	fmt.Printf("Welcome to %v\n", conference_name)
}

func getUserInput() (string, string, string, uint) {
	var firstname, lastname, email string
	var userTickets uint
	fmt.Println("Enter your first name !")
	fmt.Scan(&firstname)
	fmt.Println("Enter your last name !")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email !")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you wanna purchage !")
	fmt.Scan(&userTickets)
	return firstname, lastname, email, userTickets
}
