package main

import (
	"fmt"
	"os/exec"
)

type Client struct {
	Remote string
}

func New(remote string) (*Client, error) {
	cmd := exec.Command("lxc", "remote", "switch", remote)
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("could not switch to remote: %s", remote)
	}

	return &Client{
		Remote: remote,
	}, nil
}
