package main

import "fmt"

func main() {
	var day int
	fmt.Println("Enter the day:")
	fmt.Scanf("%d", &day)
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Sunday")

	}
}
