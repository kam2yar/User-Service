package main

import (
	"fmt"
	"github.com/kam2yar/user-service/internal"
)

func main() {
	internal.Bootstrap()
	fmt.Println("Hello world")
}
