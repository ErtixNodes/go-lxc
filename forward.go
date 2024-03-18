package lxc

import (
	"fmt"
	"os"
	"os/exec"
)

func (c *Client) ForwardPorts(name string, ports []int) error {
	for _, port := range ports {
		listen := fmt.Sprintf("listen=tcp:0.0.0.0:%d", port)
		connect := fmt.Sprintf("connect=tcp:127.0.0.1:%d", port)
		portName := fmt.Sprintf("port%d", port)
		cmd := exec.Command("lxc", "config", "device", "add", name, portName, "proxy", listen, connect)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("could not proxy port: %d vm: %s", port, name)
		}
	}

	return nil
}
