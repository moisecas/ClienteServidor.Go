package main

import (
	"fmt"
	"net"
)

func cliente() {
	c, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	ms := "saludos"
	fmt.Println(ms)
	c.Write([]byte(ms))
	c.Close()

}

func main() {
	go cliente()
	var input string
	fmt.Scanln(&input)
}
