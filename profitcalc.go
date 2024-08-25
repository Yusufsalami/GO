package main

import("fmt")



func main2(){
	revenue   := textInput("revenue: ")
	expenses  := textInput("expenses: ")
	taxRate  := textInput("tax rate: ")

	ebt , profit, ratio := calculatefinancials(revenue, expenses, taxRate )



	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	
}

func calculatefinancials(revenue,expenses,taxRate float64) (float64,float64,float64){
	ebt := revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt/profit
	return ebt,profit,ratio
}

func textInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

