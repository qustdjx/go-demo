package main

import "fmt"

func quick_sort(s []int,l int,r int)  {

	if(l<r){
		i:=l
		j:=r
		x:=s[l]

		for i<j {
			for i<j && s[j]>=x{
			  j--
			}
			if(i<j){
			 s[i]=s[j]
			 i++
			}
			for i<j && s[i]<x{
				i++
			}
			if(i<j){
				s[j]=s[i]
				j--
			}
		}
		s[i] = x;
		quick_sort(s, l, i - 1); // 递归调用
		quick_sort(s, i + 1, r);
	}

}

func main()  {
	s :=[] int {3,5,4,1,4,0,9,3,10 }
	len:=len(s)
	quick_sort(s,0,len-1)
	fmt.Println(s)
}