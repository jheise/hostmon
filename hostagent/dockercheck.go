package main

import (
	// standard
	// "fmt"
	"time"

	// local
	"github.com/jheise/hostmon/hostutil"

	// external
	"github.com/fsouza/go-dockerclient"
)

func check_docker() (*hostutil.DockerList, error) {
	dockerlist := new(hostutil.DockerList)
	// fmt.Printf("Listing Containers...\n")
	opts := docker.ListContainersOptions{true, false, 0, "", "", nil, nil}
	containers, err := client.ListContainers(opts)
	if err != nil {
		return dockerlist, err
	}

	// fill dockerlist

	for _, container := range containers {
		// fmt.Printf("container: %s\n", container)
		newcontainer := new(hostutil.Container)
		newcontainer.Name = container.Names[0]
		newcontainer.Image = container.Image
		newcontainer.Status = container.State

		// convert time to a string, so you can actually read it
		tempTime := time.Unix(container.Created, 0)
		newcontainer.Created = tempTime.Format(time.ANSIC)

		// grab all ports
		if len(container.Ports) > 0 {
			for _, port := range container.Ports {
				newnet := hostutil.ContainerNet{port.IP, port.Type, port.PublicPort}
				newcontainer.Ports = append(newcontainer.Ports, newnet)
			}
		}

		// grab all environment variables, for some reason container and apicontainer are different
		realcontainer, err := client.InspectContainer(container.ID)
		if err != nil {
			return dockerlist, err
		}
		newcontainer.Env = realcontainer.Config.Env
		dockerlist.Containers = append(dockerlist.Containers, newcontainer)
	}
	return dockerlist, err
}
