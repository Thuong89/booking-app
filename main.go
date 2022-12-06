package main

import (
	"fmt"
	"time"
)
func main() {
	//define variables
	const x = "module"
	var modulenumber int
	var username string
	var y int
	var modules_done []int  // Arrays 5elements, element data type integer. element can be left empty or some value.
	
	// ask user for their name and greet
	fmt.Print("Enter your name:")
	fmt.Scan(&username)   // & is is a pointer which came from C and C++, this to specify the memory address
	fmt.Printf("Hello %v, welcome to Go world!\n", username) 

	// pause for 3 seconds
	time.Sleep(3*time.Second)

	// introduce content, for loop
	fmt.Println("\nList",x,"to explore:")
	for modulenumber = 1;modulenumber <6; modulenumber++ {
		fmt.Println("Go",x, ":", modulenumber)
	}
	
	// ask user to select and end
	fmt.Print("Type module number you choose:")
	fmt.Scan(&y)
	fmt.Printf("\nAwesome! Enjoy your time with %v %v =)", x, y)

	// save user data and pause for 10 seconds
	modules_done = append(modules_done,y)
	time.Sleep(10*time.Second)

	// welcome back
	fmt.Printf("\n\nWelcome back %v ! Congratulations on finishing module %v! \n",username,y)
	fmt.Printf("You have learnt %v module", len(modules_done))

}