package main

import (
	"fmt"
	"sort"
)

func main()  {
	var n int
	var num int
	var s1 []int
	var s2 []int
	m := make(map[int]int)

	fmt.Scanln(&n)
	for i:=0;i<n;i++{
		var pos = 0
		fmt.Scanln(&num)
		s1 = append(s1,num)
		for _,v:=range s1{
			if v==num{
				pos++
			}
		}
		m[num]=pos
	}

	for k:=range m{
		s2 = append(s2,k)
	}
	sort.Ints(s2)
	for _,v:=range s2{
		fmt.Println(v,m[v])
	}

}