package exercises



// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    // write your code in Go 1.4
	var count int
switch N%2 {
case 0: {
	q:=N/2
	
	 count  = 1
	if q %2 ==0{

	
	for q > 0{
		if q % 2 == 0{
			count++
		}
		q/=2
	}

	}
	return count
}
default : return 0
	
}



	
}


