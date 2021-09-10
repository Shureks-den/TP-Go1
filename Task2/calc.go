package main

import (
	"fmt"
	calcModule "task2/calculator"
)

func main() {
	c := calcModule.GetPreparedData()
	res, _ := calcModule.Calculator(c)
	fmt.Println(res)
}
