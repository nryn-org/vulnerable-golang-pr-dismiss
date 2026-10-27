package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"
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
