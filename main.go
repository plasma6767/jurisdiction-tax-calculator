// Jurisdiction Tax Calculator
//
// A command-line tool that estimates personal income tax across different countries. Given a USD income, it calculates the tax owed under each supported jurisdiction's tax rules
// (progressive brackets, flat rates, zero-tax), so a user can compare take-home pay across countries.
//
// Currently supported: Singapore, UAE, Bulgaria.
package main

import (
	"fmt"
	"math"
)

type bracket struct {
	upperLimit float64
	deduction  float64
	rate       float64
	baseTax    float64
}

// dollarToPound
//
// Function added to help with the logic of the UK income tax since their are complications on the tax free income portion of the users income
func dollarToPound(dollars float64) float64 {
	dollarToPoundRate := 0.75

	pounds := dollars * dollarToPoundRate
	return pounds
}

// poundToDollar
//
// Function added to help with the logic of the UK income tax since their are complications on the tax free income portion of the users income
func poundToDollar(pounds float64) float64 {
	poundToDollarRate := 1.33

	dollars := pounds * poundToDollarRate
	return dollars
}

// dollarToEuro
//
// Function added to help with the logic of the Germany income tax for complications calculating their brackets
func dollarToEuro(dollars float64) float64 {
	dollarToEuroRate := 0.88

	euros := dollars * dollarToEuroRate
	return euros
}

// euroToDollar
//
// Funtion addedto help with the logic of the Germany income tax for complications calculating their brackets
func euroToDollar(euros float64) float64 {
	euroToDollarRate := 1.14

	dollars := euros * euroToDollarRate
	return dollars
}

// calculateSingaporeTax
//
// Calculates the amount of taxes that the user would have to pay in Singapore (in U.S. dollars)
func calculateSingaporeTax(income float64) {
	var singTax float64
	// Tax brackets in singapore (ordered by the upperlimit of the bracket), incomes get grouped into the brackets by whatever they are less than (35,000 gets grouped into the 40,000 bracket)
	// baseTax is the max tax from the previous brackets tax
	// deduction is the max of the previous bracket used to find the difference between the users income and the previous bracket, so the difference can be multiplied by the rate of that bracket
	// rate is the tax rate for that bracket, represented in decimal percentace (.1 is 10%)
	singaporeBrackets := []bracket{
		{upperLimit: 30000, deduction: 20000, rate: 0.02, baseTax: 0},
		{upperLimit: 40000, deduction: 30000, rate: 0.035, baseTax: 200},
		{upperLimit: 80000, deduction: 40000, rate: 0.07, baseTax: 550},
		{upperLimit: 120000, deduction: 80000, rate: 0.115, baseTax: 3350},
		{upperLimit: 160000, deduction: 120000, rate: 0.15, baseTax: 7950},
		{upperLimit: 200000, deduction: 160000, rate: 0.18, baseTax: 13950},
		{upperLimit: 240000, deduction: 200000, rate: 0.19, baseTax: 21150},
		{upperLimit: 280000, deduction: 240000, rate: 0.195, baseTax: 28750},
		{upperLimit: 320000, deduction: 280000, rate: 0.20, baseTax: 36550},
		{upperLimit: 500000, deduction: 320000, rate: 0.22, baseTax: 44550},
		{upperLimit: 1000000, deduction: 500000, rate: 0.23, baseTax: 84150},
		{upperLimit: math.Inf(1), deduction: 1000000, rate: 0.24, baseTax: 199150},
	}

	// looping through each bracket within sinaporeBrackets
	for _, bracket := range singaporeBrackets {
		if income < bracket.upperLimit {
			// if the users income is less than the upperLimit that we are looking at then we know that is the right bracket to look at to determine the users tax in Sinapore
			if income <= 20000 {
				// if the users income is less than $20,000 then they do not have any taxes to pay
				singTax = 0
				break
			}
			if (income - bracket.deduction) == 0 {
				// if the users income is an exact bracket number then we can just simply use the baseTax to determine their taxes because they are at the max limit of that bracket
				singTax = bracket.baseTax
				break
			}
			// The formula for their taxes is their income - the previous brackets max to get the difference, then multiply that by the rate for the piece that applies to that bracket, and add that number to the baseTax of the
			// previous bracket
			singTax = (income-bracket.deduction)*bracket.rate + bracket.baseTax
			break
		}
	}

	fmt.Printf("\nSingapore tax on your specific income: $%.2f\n\n", singTax)
}

// calculateUAETax
//
// Calculates the amount of taxes that the user would have to pay in the UAE in U.S. dollars, which is $0 for every user because the UAE has a flat 0% tax rate on all personal income
func calculateUAETax() {
	// Flat 0% tax on all personal income
	uaeTax := 0.0
	fmt.Printf("UAE tax on your specific income: $%.2f\n\n", uaeTax)
}

// calculateBulgariaTax
//
// Calculates the amount of tax that the user would have to pay in Bulgaria, which is a flat 10% rate on all personal income
func calculateBulgariaTax(income float64) {
	// Flat 10% tax rate on all personal income
	bulargiaTaxRate := 0.1
	bulgariaTax := income * bulargiaTaxRate

	fmt.Printf("Bulgaria tax on your specific income: $%.2f\n\n", bulgariaTax)
}

func calculateUSATax(income float64) {
	var usaTax float64

	// USA tax brackets, since the US uses a marginal tax system, only the upper limit of the tax bracket and the tax rate is needed
	usaBrackets := []bracket{
		{upperLimit: 124000, rate: .10},
		{upperLimit: 50400, rate: .12},
		{upperLimit: 105700, rate: .22},
		{upperLimit: 201775, rate: .24},
		{upperLimit: 256225, rate: .32},
		{upperLimit: 640600, rate: .35},
		{upperLimit: math.Inf(1), rate: .37},
	}

	// loop through the US tax brackets
	for _, bracket := range usaBrackets {
		if income <= bracket.upperLimit {
			// if the income is less than or equal to the upperlimit of the tax bracket, then it falls in that brackets tax rate
			// income multiplied by the rate gives you the amount of tax the user has to pay in the US
			usaTax = income * bracket.rate
		}
	}

	fmt.Printf("USA tax on your specific income: $%.2f\n\n", usaTax)
}

// caclulcateUKTax
//
// Caclulates the amount of taxes that the user would have to pay in the UK, using a similar type of logic to calculateSingaporeTax, only with special zeroAllowance
// for those with incomes between 100,000 and 125,140 pounds
func caclulcateUKTax(income float64) {
	pounds := dollarToPound(income)
	var ukTax float64
	var zeroAllowance float64

	if pounds < 100000 {
		zeroAllowance = 12570
	} else if pounds >= 100000 && pounds < 125140 {
		zeroAllowance = 12570 - (pounds-100000)/2
	} else {
		zeroAllowance = 0
	}

	// Brackets for the UK, works similar to how singapore brackets work, except some pieces are determined at runtime if the user has a special zeroallowance that changes their brackets
	ukBrackets := []bracket{
		{upperLimit: zeroAllowance, deduction: 0, rate: 0, baseTax: 0},
		{upperLimit: zeroAllowance + 37700, deduction: zeroAllowance, rate: 0.2, baseTax: 0},
		{upperLimit: 125140, deduction: zeroAllowance + 37700, rate: 0.4, baseTax: 7540},
		{upperLimit: math.Inf(1), deduction: 125140, rate: 0.45, baseTax: ((125140 - (zeroAllowance + 37700)) * 0.4) + 7540},
	}

	// same thing as other functions just looping through and finding when the income is less than the upperlimit of a bracket
	for _, bracket := range ukBrackets {
		if pounds <= bracket.upperLimit {
			ukTax = ((pounds - bracket.deduction) * bracket.rate) + bracket.baseTax
			break
		}
	}

	dollars := poundToDollar(ukTax)

	fmt.Printf("UK tax on your specific income: $%.2f\n", dollars)
}

func calculateGermanyTax(income float64) {
	euros := dollarToEuro(income)
	var germanyTax float64
	var scaledIncome float64 // variable to help with the scaled down income portions to calculate the specific portion of income tax

	// special income tax rule for germany require coding in else if statement for the 5 different formula possibilities
	if euros < 12349 {
		germanyTax = 0 // first zone has no income tax under 12349 euros
	} else if euros <= 17779 {
		scaledIncome = (euros - 12348) / 10000
		germanyTax = (914.51*scaledIncome + 1400) * scaledIncome // formula for the second zone of incomes
	} else if euros <= 69878 {
		scaledIncome = (euros - 17799) / 10000
		germanyTax = (173.10*scaledIncome + 2397) * 1034.87 // third zone
	} else if euros <= 277825 {
		germanyTax = 0.42*(euros-69878) + 14414.87 // fourth zone
	} else {
		germanyTax = 0.45*euros - 19470.38 // final zone
	}
}

func main() {
	var income float64

	fmt.Print("Enter your Income in US Dollars: \n$")
	_, err := fmt.Scan(&income)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	calculateSingaporeTax(income)
	calculateUAETax()
	calculateBulgariaTax(income)
	calculateUSATax(income)
	caclulcateUKTax(income)
}
