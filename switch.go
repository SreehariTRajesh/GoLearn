package main
import "fmt"
func unhex(c byte) byte{
	switch{
		case '0'<=c && c<='9':
			return 'd'
		case 'a'<=c && c<='z':
			return 'l'
		case 'A'<=c && c<='Z':
			return 'u'
	}
	return 0
}

func main(){
	fmt.Printf("%c\n",unhex('a'))
}
