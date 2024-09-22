// Boiling displays the boiling point of water
// in Celsius and Fahrenheit
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}
