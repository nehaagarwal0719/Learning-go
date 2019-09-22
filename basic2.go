package main

import "fmt"

func main(){
    
    //arrays
	var a[2]string
	a[0]="Hello"
	a[1]="World"
	fmt.Println(a)

	primes:= [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	//slices
	var s[]int = primes[2:5]
	fmt.Println(s)

	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
     //Other slices that share the same underlying array will see those changes.

     s[0]=0;
     fmt.Println(s)

     //A slice literal is like an array literal without the length.
     //different ways of slicing 
    var s1[]int=primes[:4]
    fmt.Println(s1)
    
    var s2[]int=primes[:]
    fmt.Println(s2)

    var s3[]int=primes[0:]
    fmt.Println(s3)

    //The length of a slice is the number of elements it contains.
    //The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
    fmt.Printf("len=%d cap=%d \n",len(s3),cap(s3))


    //nill slice has length and capacity zero
    if s == nil {
		fmt.Println("nil!")
	}
    

    //The make function allocates a zeroed array and returns a slice that refers to that array
    a1 := make([]int, 5) 
    fmt.Println(a1)
    //to define capacity also
    b := make([]int, 0, 5) // len(b)=0, cap(b)=5


}