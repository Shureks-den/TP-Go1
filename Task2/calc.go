package main

import (
	"fmt"
	calcModule "task2/calculator"
)

func main() {
	c := calcModule.GetPreparedData()
	res, err := calcModule.Calculator(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
