package lxc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func setPassword(name string, password string) error {
	passwdInit := fmt.Sprintf("echo -e '%s\n%s' | passwd root", password, password)
	cmd := exec.Command("lxc", "exec", name, "--", "bash", "-c", passwdInit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not setup passwd vm: %s", name)
	}

	return nil
}

func removeDefaultUser(name string) error {
	cmd := exec.Command("lxc", "exec", name, "--", "bash", "-c", "rm -rf /home/ubuntu")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not delete default user vm: %s", name)
	}

	return nil
}

func initSSH(name string, port int) error {
	cmds := []string{
		"echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config",
		"echo 'IgnoreUserKnownHosts no' >> /etc/ssh/sshd_config",
		"echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config",
		"echo 'PubkeyAuthentication yes' >> /etc/ssh/sshd_config",
		"sed -i '/UsePAM yes/c\\UsePAM no' /etc/ssh/sshd_config",
		fmt.Sprintf("echo 'Port %d' >> /etc/ssh/sshd_config", port),
		"ssh-keygen -q -t rsa -N '' -f ~/.ssh/id_rsa <<<y >/dev/null 2>&1",
		"ssh-add",
		"service ssh restart",
		"service sshd restart",
	}
	sshInit := strings.Join(cmds, " && ")
	cmd := exec.Command("lxc", "exec", name, "--", "bash", "-c", sshInit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not delete default user vm: %s", name)
	}

	return nil
}
