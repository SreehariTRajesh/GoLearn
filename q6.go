package main
func f(x int,y int) (int,int) {
	if(x>y){
		return y,x
	}
	return x,y
}
func main(){
	x,y:=10,12;
	println(x,",",y)
	x,y=f(x,y)
	println(x,",",y)
}
