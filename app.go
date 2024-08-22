package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
