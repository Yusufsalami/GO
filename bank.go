package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balane.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil{
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance , err := strconv.ParseFloat(balanceText,64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance")
	}
	return balance , nil
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText) , 0644)
}

func main3(){
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		fmt.Println("-------")

	}

	fmt.Println("Welcome to Go Bank!")

	for {
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2{
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Your new balance is", accountBalance)
		writeBalanceToFile(accountBalance)
	}else if choice ==3 {
		fmt.Print("How much do you wish to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}
		if withdrawAmount > accountBalance{
			fmt.Println("Insufficient funds")
			return
		}

		fmt.Println("you withdrew",withdrawAmount)
		accountBalance -= withdrawAmount
		
		fmt.Println("Your current balance is", accountBalance)
	}else {
		fmt.Println("Goodbye!")
	}

	fmt.Println("Your  choice") 
    }

}