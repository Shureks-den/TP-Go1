package main

import (
	"log"
	uniqueModule "task1/uniq"
)

func main() {
	var data []string
	var fl uniqueModule.Flags
	uniqueModule.ParseFlags(&fl)
	if err := uniqueModule.Read(&data, &fl); err != nil {
		log.Fatal(err)
	}
	if err := uniqueModule.Write(data, &fl); err != nil {
		log.Fatal(err)
	}
}
