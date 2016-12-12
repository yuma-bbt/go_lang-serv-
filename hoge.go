package main


import (
	"fmt"
)

func main() {
	var foo int =3 
	var ans  int
	for  i :=0;  i < 100000000000000 ; 
	{
		if  (i % 3) == 0 {
			ans = foo * i 
			fmt.Println(ans)
			i++
		}
		i++
		//time.Sleep(100*time.Millisecond)
	}
}
