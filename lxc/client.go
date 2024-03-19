package lxc

import (
	"fmt"
	"os/exec"
)

type Client struct {
	Remote string
	Host   string
}

func New(remote string, ip string) (*Client, error) {
	cmd := exec.Command("lxc", "remote", "switch", remote)
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("could not switch to remote: %s", remote)
	}

	return &Client{
		Remote: remote,
		Host:   ip,
	}, nil
}
