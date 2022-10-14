package main

import (
	"fmt"
)

func Num(name string)int {
	Map:=map[string]int{"A":1, "B":2, "C":3, "D":4, "E":5, "F":6, "G":7, "H":8,"I":9,"J":10,"K":11,"L":12,"M":13, "N":14, "O":15,"P":16,"Q":17,"R":18,"S":19,"T":20,"U":21,"V":22,"W":23,"X":24,"Y":25,"Z":26}
	var sum int=1
	for _,v:= range name{
		sum=sum*(Map[string(v)])
	}
	return sum%47
}
func main()  {
	var (
		name1 string
		name2 string
	)
	fmt.Scan(&name1)
	fmt.Scan(&name2)
	if Num(name1)==Num(name2){
		fmt.Println("GO")
	} else{
		fmt.Println("STAY")
	}

}

