// ship-server
package main

//go:generate protoc3 -I $GOPATH/src/github.com/LSFN/seprotocol/ --go_out=. $GOPATH/src/github.com/LSFN/seprotocol/upstream.proto $GOPATH/src/github.com/LSFN/seprotocol/downstream.proto

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
