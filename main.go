package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Package level variable. Cannot be declared using Syntax Sugar
// Accesible everywhere
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)
var bookingStruct = make([]User, 0)

// The main goroutine does not wait for other routines to finish their execution.
// WaitGroup makes main goroutine wait for the execution of other goroutines
// It has 3 methods
// Add, Wait and Done.
var wg = sync.WaitGroup{}

// Struct can have entities of different data types. Map has entities of same data type
type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// Alternate syntax (Syntax sugar)
	// conferenceName := "Go Conference"
	// Since value is assigned, go automatically infers the data type
	// var conferenceName = "Go Conference"
	// const conferenceTickets = 50
	// uint = unsigned integer. Only positive whole numbers
	// var remainingTickets uint = 50

	// var bookings []string
	// if array
	// var bookings [50]string

	// fmt.Printf("conferenceName is of type %T, remainingTickets is of type %T, conferenceTickets is of type %T\n", conferenceName, remainingTickets, conferenceTickets)

	// %v and %T used in Printf function are used to format the output. The are called placeholder or annotation verbs.
	// %v stands for value. %T stands for type

	greetUser()

	// In Go, a function can have multiple return values
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if !isValidName || !isValidEmail || !isValidTicketNumber {
		fmt.Println("Please enter valid input!!")
		// continue
	}

	bookTickets(userTickets, firstName, lastName, email)
	// The function will run in a new thread
	// go creates a new go routine
	// goroutine is a lightweight thread managed by go runtime

	wg.Add(1)
	// Sets the number of goroutines to wait for
	// increases the counter by the provided number
	go sendTicket(userTickets, firstName, lastName, email)

	firstNames := getFirstNames(bookings)

	fmt.Printf("The first names of bookings are: %v\n", firstNames)
	fmt.Printf("The list of bookings Map is: %v\n", bookings)
	fmt.Printf("The list of bookings Struct is: %v\n", bookingStruct)

	// fmt.Printf("Print the slice: %v\n", bookings)
	// fmt.Printf("Print the first index of bookings slice: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Size of Slice bookings: %v\n", len(bookings))

	if remainingTickets == 0 {
		fmt.Println("Booking full. Come back next year")
		// break
	}

	// Blocks until the waitgroup counter is 0
	wg.Wait()

	// SWITCH STATEMENT
	// city := "Pune"

	// switch city {
	// case "Pune":
	// 	// some code of Pune conference tickets
	// case "Aurangabad", "Jalna":
	// 	// some code of A'bad and Jalna conference tickets
	// case "Nagpur", "Chandrapur":
	// 	// some code of Nagpur and Chandrapur conference tickets
	// default:
	// 	fmt.Println("No valid city selected")
	// }

}

func greetUser() {
	fmt.Printf("Welome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets are %v and tickets remaining are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!!")
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userStruct = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Map of Users
	var userData = make(map[string]string)
	userData["firstname"] = firstName
	userData["lastname"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	bookingStruct = append(bookingStruct, userStruct)

	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// Fields spilts a string with whitespace as a seperator
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking["firstname"])
	}
	// If array
	// bookings[0] = firstName + " " + lastName
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Pointers are special variables that points to the memory address of another variable
	// fmt.Println(firstName)
	// fmt.Println(&firstName)

	fmt.Println("Enter your firstName")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// delay the execution by 10 seconds
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("##################")
	// Decrements the waitgroup counter by 1
	wg.Done()
}
