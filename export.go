package main

import (
	"fmt"
	"os/exec"
)

func (c *Client) CreateFromExport(inputPath string) error {
	cmd := exec.Command("lxc", "import", inputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not initialize vm err: %v", err)
	}

	return nil
}

func (c *Client) ExportServer(name string, outputPath string) error {
	cmd := exec.Command("lxc", "export", name, outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not backup vm err: %v", err)
	}

	return nil
}
