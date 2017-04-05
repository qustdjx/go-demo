package main

import "fmt"

func sum(n uint64) uint64 {

	var res uint64
	res=0
        var i uint64
	for i=1;i<=n;i++ {
	 res=i+res
	}
	return res
}

func main()  {
	var num uint64
	num=10000000000
	fmt.Println(sum(num))

}
