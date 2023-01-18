package main
func main(){
	var p *int
	var i int
	p=&i
	*p=8
	println("&i:",p);
	println("i",i);
	println("*p",*p);
}
