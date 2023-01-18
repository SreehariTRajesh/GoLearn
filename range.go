package main
func main(){
	list:=[]string{"a","b","c","d","e","f"}
	for k,v:= range list{
		println("{%d:%s}",k,v);
	}
}
