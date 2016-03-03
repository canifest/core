package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benschw/opin-go/rest"
	"github.com/canifest/core/dockerfile/api"
)

var _ = log.Print

type DockerfileClient struct {
	Host string
}

func (c *DockerfileClient) AddDockerfile(dockerfileJSON api.Dockerfile) (api.Dockerfile, error) {
	var dockerfile api.Dockerfile

	url := fmt.Sprintf("%s/api/dockerfile", c.Host)
	fmt.Printf("Making Post to %s", url)
	r, err := rest.MakeRequest("POST", url, dockerfileJSON)
	if err != nil {
		return dockerfile, err
	}
	err = rest.ProcessResponseEntity(r, &dockerfile, http.StatusCreated)
	return dockerfile, err
}

func (c *DockerfileClient) FindDockerfile(id int) (api.Dockerfile, error) {
	var dockerfile api.Dockerfile

	url := fmt.Sprintf("%s/api/dockerfile/%d", c.Host, id)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return dockerfile, err
	}
	err = rest.ProcessResponseEntity(r, &dockerfile, http.StatusOK)
	return dockerfile, err
}
