package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount) // se le pasa un pointer de la variable donde se guardara la entrada del raton
	outputText("years to invest: ")
	fmt.Scan(&years)
	fmt.Print("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)
	//var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %v\n", futureValue)
	formattedRFV := fmt.Sprint("Future Value (adjusted for inflation): %0f", futureRealValue)
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Printf("Future Value: %v\n", futureValue)
	//fmt.Printf("Future Value (adjusted for inflation): %0f", futureRealValue) // el numero a la izq de la f es el num de decimales del float que se imprimira

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
