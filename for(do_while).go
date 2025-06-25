package main
import "fmt"
func main() {
	var temp, sum int
	i := 0
	for {
		if i >= 5 {
			break
		}
		fmt.Scanf("%d\n", &temp)
		sum = sum + temp
		i++
	}
	fmt.Printf("%d", sum)
}