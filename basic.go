package main

import(
	  "fmt"
)

//availabe throught the package
const(
	e="neha"
	//declared within this const scope
	i=iota
	j=iota
) 

const(
	//declared with this const scope
	k=iota
)

type user struct{
	fname string
	lname string
	age int
}

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


    //IOTA
    //Can use it as counter variable
    fmt.Println(i,j)
    fmt.Println(k)


    //POINTERS
    //prints address of a
    fmt.Println(&a)

    var g *int
    g=&a
    fmt.Println(g,*g)


    //STRUCTURES
    //1 way
    student := user{}
    student.fname="neha"
    student.lname="agarwal"
    student.age=21
    
    //2 way
    student1 := user {
        fname:"Payal",
        lname:"Agarwal",
        age:25,
    }

    fmt.Println(student)
    fmt.Println(student1)


    //ARRAYS
    //length needs to be defined
    var arr [5]int
  


}

//if function name starts with capital letter it can be accessible outside the package 
//if it starts with lowercase letter then it will be private to the package

//func Cal(a, b int) int --if both parameters are of same type
//func Cal(a int , b int) (int,error) --can have multiple return types -- error is a golang library that returns a string
//func Cal(a, b int) int --if both parameters are of same type
//func Cal(a int , b int, c=10) (int,error) 
func Cal(a int , b int)  (x int) { //if you just want to write return then specify variable name with return then you dont have to declare x
 
  x=a+b
   return 
}


