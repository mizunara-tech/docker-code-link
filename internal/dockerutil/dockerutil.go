package dockerutil

import (
	docker "github.com/fsouza/go-dockerclient"
)

func GetContainerWorkingDir(client *docker.Client, containerName string) (string, error) {
	container, err := client.InspectContainerWithOptions(docker.InspectContainerOptions{ID: containerName})
	if err != nil {
		return "", err
	}
	return container.Config.WorkingDir, nil
}
