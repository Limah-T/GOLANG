package main
import ("fmt")
const NUMERATOR int = 5
const DENOMINATOR int = 9

func temperatureConverter(converter string, temperature int){
	if converter == "C" {
		var result = (temperature * DENOMINATOR/NUMERATOR) + 32
		fmt.Printf("Temperature is %v°F", result)
	} else if converter == "F" {
		var result = (temperature - 32) * NUMERATOR/DENOMINATOR
		fmt.Printf("Temperature is %v°C", result)
	} else {
		fmt.Print("Invalid converter input")
	} 
}
func converter() {
	var converter string
	var temperature int
	fmt.Print("Enter the Converter 'F' for Fahrenheit, 'C' for Celsuis: ")
	fmt.Scan(&converter)
	fmt.Print("What is the temperature? ")
	fmt.Scan(&temperature)
	temperatureConverter(converter, temperature)

}
