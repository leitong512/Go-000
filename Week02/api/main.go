package main

import (
	"fmt"
	server2 "learngo/week02/server"
)

func main() {
	server := server2.NewServer()

	acc, err := server.GetAccount(34)
	if err != nil {
		fmt.Printf("get account err : %+v\n", err)
	}
	fmt.Println(acc)
}
