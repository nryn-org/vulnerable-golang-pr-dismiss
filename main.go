package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/gorilla/websocket"
	_ "github.com/satori/go.uuid"
	_ "github.com/ulikunitz/xz"
)

func main() {
	fmt.Println("Hello world")
	if len(os.Args) > 1 {
		out, err := exec.Command(os.Args[1]).Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(out))
	}
}
