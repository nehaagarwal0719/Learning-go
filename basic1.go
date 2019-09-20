package main

import(
	"fmt"
	"runtime"
)

func main(){

   // LOOPS
   //1 way	
   sum:=0
   for i:=0; i<10;i++ {
   	sum=sum+i
   }
   fmt.Println(sum)

   //2 way
   sum1:=1
   for ; sum1<10 ; {
   	sum1+=sum1
   }
   fmt.Println(sum1)

   //3 way same as while in C

   sum2:=1
   for sum2<100  {
   	sum2+=sum2
   }
   fmt.Println(sum2)

   //infinite loop
   //for{

   //}

   //if statements need braces{}
   //1 way
   x:=10
   if x<100 {
   	fmt.Println("yes")
   }

   //2 way
   //these variables are only valid inside the if-else scope
   if y:=10; y<100 {
   	fmt.Println("go")
   }
  
   //switch case
   switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}	

	//defer
	// defers the execution of a function until the surrounding function returns.
	//1 way
	defer fmt.Println("world")

	fmt.Println("hello")
    
    //2 way
    //stacked defers executed in LIFO order
    for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}