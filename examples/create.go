package main

import (
	"fmt"
	"strings"

	"github.com/ErtixNodes/go-lxc/lxc"

	"github.com/google/uuid"
)

func main() {
	client, err := lxc.New("vps", "127.0.0.1")
	if err != nil {
		panic(err)
	}

	randomName := strings.Split(uuid.NewString(), "-")[0]
	err = client.CreateServer(randomName, &lxc.CreateServerParams{
		Memory:  2048,
		Storage: 8000,
		Cores:   2,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Created server with name: %s\n", randomName)
	err = client.ForwardPorts(randomName, []int{2022, 2023})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Installing server...")
	err = client.InstallServer(randomName, "1234", 2022)
	if err != nil {
		fmt.Println(err)
	}

}
