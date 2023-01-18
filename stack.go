package main
import "fmt"
type stack struct{
	i int
	data [10] int
}

func (s stack) push(k int){
	if(s.i+1>9){
		println("Stack is full");
	}else{
		s.data[s.i]=k;
		s.i++;
	}
}
func (s *stack) pop() int{
	if(s.i==0){
		println("Stack is empty")
		return 0;
	}
	s.i--;
	return s.data[s.i];
}
func main(){
	var s stack;
	s.push(25)
	s.push(14)
	fmt.Printf("%v",s);
}
