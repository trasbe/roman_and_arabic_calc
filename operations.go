/*
package main

import "strings"

type MyFloat float64

	func (f MyFloat) Abs() float64 {
		if f < 0 {
			return float64(-f)
		}
		return float64(f)
	}

	func main() {
		f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Abs())
	}

	func main() {
		S := "Hello World"
		println(S)
		println(strings.Contains(S, "Hell"))
	}
*/
/*
package main

import (
	"fmt"
)

func add_(x float32, y float32) {
	var res float32 = x + y
	fmt.Printf("result is : %.f\n", res)
}

func substraction_(x float32, y float32) {
	var res float32 = x - y
	fmt.Printf("result is : %.f\n", res)
}

func divison_(x float32, y float32) {
	var res float32 = x / y
	fmt.Printf("result is : %f\n", res)
}
func multiplication_(x float32, y float32) {
	var res float32 = x * y
	fmt.Printf("result is : %f\n", res)
}
*/
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main_() {
	var input string
	fmt.Println("Enter a Roman numeral between 1 and 10: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	// Validate the input using a regular expression
	valid, _ := regexp.MatchString(`^(I|II|III|IV|V|VI|VII|VIII|IX|X)$`, input)
	if valid {
		fmt.Println("Valid Roman numeral:", input)
	} else {
		fmt.Println("Invalid Roman numeral. Please enter a Roman numeral between 1 and 10.")
	}
}
