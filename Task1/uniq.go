package main

import (
	uniqueModule "task1/uniq"
)

func main() {
	var data []string
	var fl uniqueModule.Flags
	uniqueModule.ParseFlags(&fl)
	uniqueModule.Read(&data, &fl)
	uniqueModule.Write(data, &fl)
}
