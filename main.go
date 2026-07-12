package main

import (
	"fmt"
	"math"
)

type country struct {
	Name    string
	taxRate map[string]int
}

func calculateSingaporeTax(income float64) {
	var singTax float64
	// extra list to loop through map keys in correct order
	singTaxKeys := []float64{30000, 40000, 80000, 120000, 160000, 200000, 240000, 280000, 320000, 500000, 1000000, math.Inf(1)}
	// Mapping the tax bracket, to a list of the deduction portion, the tax rate for that deduction portion, and the gross tax payable for that bracket
	singTaxRate := map[float64][]float64{
		30000:       {20000, 0.02, 0},
		40000:       {30000, 0.035, 200},
		80000:       {40000, 0.07, 550},
		120000:      {80000, 0.115, 3350},
		160000:      {120000, 0.15, 7950},
		200000:      {160000, 0.18, 13950},
		240000:      {200000, 0.19, 21150},
		280000:      {240000, 0.195, 28750},
		320000:      {280000, 0.20, 36550},
		500000:      {320000, 0.22, 44550},
		1000000:     {500000, 0.23, 84150},
		math.Inf(1): {1000000, 0.24, 199150},
	}

	for _, value := range singTaxKeys {
		if income < value {
			if income <= 20000 {
				singTax = 0
				break
			}
			if (income - singTaxRate[value][0]) == 0 {
				singTax = singTaxRate[value][2]
				break
			}
			singTax = (income-singTaxRate[value][0])*singTaxRate[value][1] + singTaxRate[value][2]
			break
		}
	}

	fmt.Printf("Singapore tax on your specific income: %.2f\n", singTax)
}

func calculateUAETax() {
	// Flat 0% tax on all personal income
	uaeTax := 0.0
	fmt.Printf("UAE tax on your specific income: %.2f\n", uaeTax)
}

func calculateBulgariaTax(income float64) {
	// Flat 10% tax rate on all personal income
	bulargiaTaxRate := 0.1
	bulgariaTax := income * bulargiaTaxRate

	fmt.Printf("Bulgaria tax on your specific income: %.2f\n", bulgariaTax)
}

func main() {
	var income float64

	fmt.Print("Enter your Income in US Dollars: ")
	_, err := fmt.Scan(&income)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Printf("You make %.2f US Dollars per year.\n", income)
	calculateSingaporeTax(income)
	calculateUAETax()
	calculateBulgariaTax(income)
}
