package main

import (
	"fmt"
)

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	// text := "The following isthe account information."
	// first_name := "Luke"
	last_name := "Skywalker"
	age := 20 // int
	// weight := 73.0 + "kg"  // concat will fail - mismatch types
	// height := 1.72 + "m"   // concat will fail - mismatch types
	remaining_credits := "123.55" //
	currency_symbol := '$'        // this should be a rune as it is declared in single characters
	// account_name := "admin"
	// account_password := "password"
	// subscribed_user := "true"

	fmt.Printf("%T %T %T %T\n", last_name, age, currency_symbol, remaining_credits)
	fmt.Printf("The Unicode CodePoint for currency symbol is %U\n: ", currency_symbol)
	//The decimal value of Unicode Code Point for currency symbol is
	fmt.Println("..whose decimal value is: ", rune(currency_symbol))
}

func main() {
	tellMeTypes()
}
