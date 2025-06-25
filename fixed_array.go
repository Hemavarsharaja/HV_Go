package main
import "fmt"
func main() {
	var arr [5]int
	arr=[...]int{1, 2, 3, 4, 5}
	for i:=0;i<len(arr);i++ {
		fmt.Printf("%d: %d\n", i,arr[i])
	}
}