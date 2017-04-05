package main

import "fmt"

const TEN_MILION  uint64 = 10000000

func sum(n uint64) uint64  {
	if(n<2){
	  return 1
	}
	return n+sum(n-1)
}

func main()  {
	var num uint64
	num=1000000000
	fmt.Println(sum(num))
}