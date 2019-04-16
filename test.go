package main

import (
	"fmt"
)

func main(){
	var a, b, c = getStr("zy")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(getStr("zy"))
}

func getStr(a string) (string,string,int8){
	return a + " is an SB" , a + " HaHa", 11
}
