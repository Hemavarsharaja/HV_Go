package main
import "fmt"
func main(){
	var a int;
	fmt.Println("Enter a number")
	fmt.Scanf("%d",&a)
	if a%2==0{
		fmt.Printf("%d is Even",a)
	}else{
		fmt.Printf("%d is Odd",a)
	}
}