package main

import "fmt"

// if-else statement
func main() {
	/*fmt.Print("Enter a number :")
		var input int
		fmt.Scanln(&input)

		if input%2 == 0 {
			fmt.Println("number is even")
		} else {
			fmt.Println("number is odd")
		}
	}*/
	// switch case
	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 101
		fmt.Println(data)
		fallthrough
	case 11:
		data = 102
		fmt.Println(data)
		fallthrough
	case 15:
		data = 103
		fmt.Println(data)
	default:
		data = 10010
		fmt.Println(data)
	}
}
