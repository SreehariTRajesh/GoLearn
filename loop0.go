package main
import "fmt"
func main(){
	a:=[]int {1,2,3,4,5,6,7,8,9}
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1 {
		a[i],a[j]=a[j],a[i]
	}
	for i:=0;i<9;i++{
		fmt.Printf("%d\n",a[i])
	}
}
