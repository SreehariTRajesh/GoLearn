package main
func main(){
	a:=[]int{};
	for i:=0;i<10;i++ {
		a=append(a,i)
	}
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1 {
		b:=a[i:j]
		for k:=0;k<len(b);k++ {
			println(a[k])
		}
	}
}
