package main
import "fmt"
import "os"
func dosomething(f){

}
func main(){
	name:="hello.go"
	f,err:=os.Open(name,os.O_RDONLY,0)
	if err!=nil{
		fmt.Printf("ERROR:%v",err)
	}
	dosomething(f)
}

