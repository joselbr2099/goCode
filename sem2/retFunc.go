package main

import(
	"fmt"
)


//declare funcs
func incFn(x int) int {
return x+1
}

//declare in one line
func decFn(x int) int {return x-1}


func main(){

	fmt.Printf("%d",applyIt(decFn,5))
	fmt.Printf("%d",testing(3))

	fmt.Println("\nreturn function")
	sum:=retSum(4)
	fmt.Println(sum(1))
	
}

//functions takes a function as argument and return a function
func applyIt(varFunc func(int) int, val int) int{
	return varFunc(val)
}


func testing(x int)int{
	return decFn(x)
}

//function can return a function

func retSum(x int) func (int) int {
	fn:= func (int) int{
		return x+1
	}

	return fn
}

