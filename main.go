package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var running = 99

	for running == 99 {
		var decision int
		fmt.Println("Welcome to SimpleBMICalc!")
		time.Sleep(2 * time.Second)
		fmt.Println("What would you like to do?")
		fmt.Println("Calculate Height in Inches (1) \nCalculate BMI (2)\nExit Program (3)")

		_, err := fmt.Scanf("%d", &decision)

		switch decision {
		case 1:
			inputInches()
			break
		case 2:
			inputBMI()
			break
		case 3:
			running = decision
			break
		default:
			println("Please enter a number!")
		}

		if err != nil {
			fmt.Print("We've ran into an error: ", err)
		}
	}

}

func inputInches() {
	var feet int
	var inches int

	fmt.Println("Please enter your height in feet (The 5 in 5' 10\") ")
	_, err := fmt.Scanf("%d", &feet)

	fmt.Println("Please enter your height in inches (The 10 in 5' 10\") ")
	_, errord := fmt.Scanf("%d", &inches)

	calcInches(feet, inches, err, errord)
}

func inputBMI() {
	var inches float64
	var weight float64

	fmt.Println("Please enter your height in inches \nIf you don't know your height in inches, go back and calculate!")
	_, errInches := fmt.Scanf("%f", &inches)

	fmt.Println("Please enter your weight in pounds!")
	_, errWeight := fmt.Scanf("%f", &weight)

	if errInches != nil {
		fmt.Println("FATAL ERROR: ", errInches)
	}

	if errWeight != nil {
		fmt.Println("FATAL ERROR :", errWeight)
	}

	calcBMI(inches, weight)
}

func calcInches(feet int, inches int, err error, errord error) {

	feet = feet * 12

	var final = feet + inches

	fmt.Println(final)

	if err != nil {
		fmt.Println("FATAL ERROR:", err)
	}

	if errord != nil {
		fmt.Println("FATAL ERROR:", err)
	}

}

func calcBMI(inches float64, weight float64) {

	//IMPERIAL CALCULATIONS
	var BMI float64

	weight = weight * 703
	inches = math.Pow(inches, 2)

	BMI = Round(weight/inches, .5, 2)

	switch {
	case BMI < 18.5:
		fmt.Println("Your BMI is:", BMI, "- which is Underweight!")
		break
	case BMI > 18.5 && BMI < 24.9:
		fmt.Println("Your BMI is: ", BMI, "- which is Normal!")
		break
	case BMI > 25.0 && BMI < 29.9:
		fmt.Println("Your BMI is: ", BMI, "- which is Overweight!")
		break
	default:
		fmt.Println("Your BMI is: ", BMI, "- which is Obese!")
		break
	}
}

// Round : Rounds the supplied value, and returns the rounded value.
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
