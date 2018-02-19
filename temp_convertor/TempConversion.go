//TempConversion prints a table of temperature conversion from 0 to 100 degrees Celsius
package main

import "fmt"

//CtoF converts Celsius to Fahrenheit
func CtoF(celsius float64) float64 {
	fahr := ((9.0 / 5.0) * celsius) + 32
	return fahr
}

func main() {
	fmt.Println("Celsius | Fahrenheit")

	for tempC := 0.0; tempC <= 100; tempC++ {
		tempF := CtoF(tempC)
		fmt.Printf("%.1f | %.1f\n", tempC, tempF)
	}
}
