package main
import "fmt"

func main(){
	fmt.Printf("If statements in Go\n");
	var x int
	x=10; 
	if x>9 {
		fmt.Printf("Inside If\n")
	} else{
		fmt.Printf("Inside Else\n")
	}
	if true && true {
		fmt.Printf("true && true\n")
	}
	if !false {
		fmt.Printf("!false\n");
	}
}
