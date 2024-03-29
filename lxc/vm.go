package lxc

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func (c *Client) ListServers() (*[]Server, error) {
	cmd := exec.Command("lxc", "list", "--format", "yaml")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not list vms err: %v", err)
	}

	servers := []Server{}
	err = yaml.Unmarshal(output, &servers)
	if err != nil {
		return nil, err
	}

	return &servers, nil
}

func (c *Client) GetServer(name string) (*ServerInfo, error) {
	cmd := exec.Command("lxc", "info", name)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("vm not found err: %v", err)
	}

	server := ServerInfo{}
	err = yaml.Unmarshal(output, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (c *Client) DeleteServer(name string) error {
	cmd := exec.Command("lxc", "delete", name, "--force")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could nog delete vm err: %v", err)
	}

	return nil
}

func (c *Client) CreateServer(name string, options *CreateServerParams) error {
	cmd := exec.Command("lxc", "launch", "ubuntu", name, "-c", fmt.Sprintf("limits.cpu=%d", options.Cores), "-c", fmt.Sprintf("limits.memory=%dMiB", options.Memory))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could nog create vm err: %v", err)
	}

	cmd = exec.Command("lxc", "config", "device", "override", name, "root", fmt.Sprintf("size=%dMB", options.Storage))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("could nog set storage limit for vm err: %v", err)
	}

	return nil
}

func (c *Client) StartServer(name string) error {
	cmd := exec.Command("lxc", "start", name)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could nog start vm err: %v", err)
	}

	return nil
}
