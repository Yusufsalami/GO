package main

import (
	"fmt"
	"math"
)

 const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var expectedReturnRate float64 = 5.5
	var years float64

	//fmt.Print("investment Amount: ")
	outputText("investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("years: ")
	outputText("years: ")
	fmt.Scan(&years)

	//fmt.Print("expected Return Rate: ")
	outputText("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	//fmt.Printf("Future value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue,futureRealValue)
	//fmt.Println(futureRealValue)

	print(formattedFV,formattedFRV)
}

func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64,float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}