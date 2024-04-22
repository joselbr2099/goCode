package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


//function to sort user slice
func BubbleSort(userSlice *[]int){

	//changues var
	asSwap:=false

	//len of SLice
	lenSlice:=len(*userSlice)
	//fmt.Println("lenS",lenSlice)
	//main loop
	for {

		//lop for userSlice
		for index:=0;index<lenSlice-1; index++{
			if (*userSlice)[index]>(*userSlice)[index+1]{
				Swap(userSlice,index)
				asSwap=true
			}
		}

		//if end sort
		if asSwap{
			asSwap=false
			fmt.Println("--> new round")
		}else{
			break
		}
	}
}

//swap silce values
func Swap(userSlice *[]int, index int){
	tmp:=(*userSlice)[index]
	(*userSlice)[index]=(*userSlice)[index+1]
	(*userSlice)[index+1]=tmp
	fmt.Println("swaping",*userSlice)
}


//function to get error
func getErr(err error){
	if err!=nil{
		fmt.Println("Error: ",err)
		os.Exit(3)
	}
}


//function to prepare and insert each number to slice fro user
func main(){

	//slice of user input
	userNumbers:=make([]int,0)


	//read numbers from user
	fmt.Println("insert numbers separate with spaces:")
	scan := bufio.NewReader(os.Stdin)
	tmpNumbers,err:=scan.ReadString('\n') //read input to \n
	tmpNumbers=strings.TrimSuffix(tmpNumbers,"\n") //remove \n from end of imput
	tmpNumbers=strings.TrimSuffix(tmpNumbers," ") //remove space from end imput (if exist)
	getErr(err) //check for errors

	//insert numbers in slice
	for _,number:=range strings.Split(tmpNumbers, " "){  //extract all numbers from tmpNumbers
		tmp,err:=strconv.Atoi(number)		//convert each number to int
		getErr(err)
		userNumbers=append(userNumbers,tmp) //insert each number to slice
	}

	//call to BubleSort function
	BubbleSort(&userNumbers)

	fmt.Println("sorted slice: ",userNumbers)
    fmt.Println("sorte  slice")
}
