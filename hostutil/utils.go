package hostutil

type ContainerNet struct {
	IPAddr  string `json:"IPAddr"`
	NetType string `json:"NetType"`
	Port    int64  `json:"Port"`
}

type Container struct {
	Name    string         `json:"Name"`
	Image   string         `json:"Image"`
	Status  string         `json:"Status"`
	Created string         `json:"Created"`
	Ports   []ContainerNet `json:"Ports"`
	Env     []string       `json:"Env"`
}

type DockerList struct {
	Containers []*Container `json:"Containers"`
}

type Host struct {
	Hostname   string       `json:"Hostname"`
	Containers []*Container `json:"Containers"`
}
