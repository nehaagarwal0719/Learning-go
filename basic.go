package main

import(
	  "fmt"
)

//availabe throught the package
const(
	e="neha"
) 
func main(){
	//printing
	fmt.Println("Hello everyone")

	//VARIABLES
	//variable declaration 1 way
	var a int
	a=9
	fmt.Println(a)

	//variable declaration 2 way
	var b int =10
	fmt.Println(b)

	//variable declaration 3 way
	c:="neha"
	fmt.Println(c)

	//CONSTANTS
	//value of d cannot be changed
	// it can only be declared in this way
	//availabe only inside this function
    const d =2.1

}

//if function name starts with capital letter it can be accessible outside the package 
//if it starts with lowercase letter then it will be private to the package
/*
func Cal(a int, b int){ 

}*/
