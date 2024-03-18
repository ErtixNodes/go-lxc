package lxc

import "time"

type ServerInfo struct {
	Resources Resources `yaml:"Resources"`
}

type Remote struct {
	Addr   string `yaml:"addr"`
	Public bool   `yaml:"public"`
}

type Server struct {
	Name         string    `yaml:"name"`
	Description  string    `yaml:"description"`
	Status       string    `yaml:"status"`
	StatusCode   int       `yaml:"status_code"`
	CreatedAt    time.Time `yaml:"created_at"`
	LastUsedAt   time.Time `yaml:"last_used_at"`
	Location     string    `yaml:"location"`
	Type         string    `yaml:"type"`
	Project      string    `yaml:"project"`
	Architecture string    `yaml:"architecture"`
	Ephemeral    bool      `yaml:"ephemeral"`
	Stateful     bool      `yaml:"stateful"`
}

type CreateServerParams struct {
	Memory         int
	Storage        int
	Cores          int
	ForwardedPorts []string
}

type Resources struct {
	Processes   int                `yaml:"Processes"`
	DiskUsage   DiskUsage          `yaml:"Disk usage"`
	CPUUsage    CPUUsage           `yaml:"CPU usage"`
	MemoryUsage MemoryUsage        `yaml:"Memory usage"`
	Network     map[string]Network `yaml:"Network usage"`
}

type DiskUsage struct {
	Root string `yaml:"root"`
}

type CPUUsage struct {
	Usage string `yaml:"CPU usage (in seconds)"`
}

type MemoryUsage struct {
	Current string `yaml:"Memory (current)"`
}

type Network struct {
	Type            string            `yaml:"Type"`
	State           string            `yaml:"State"`
	HostInterface   string            `yaml:"Host interface"`
	MACAddress      string            `yaml:"MAC address"`
	MTU             int               `yaml:"MTU"`
	BytesReceived   string            `yaml:"Bytes received"`
	BytesSent       string            `yaml:"Bytes sent"`
	PacketsReceived int               `yaml:"Packets received"`
	PacketsSent     int               `yaml:"Packets sent"`
	IPAddresses     map[string]string `yaml:"IP addresses"`
}
