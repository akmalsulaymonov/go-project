package main

import (
	"fmt"

	"github.com/akmalsulaymonov/go-project/internal/network"
)

func main() {
	fmt.Println("Awesome CLI v.0.0.1")

	network.Ping("127.0.0.1")
}
