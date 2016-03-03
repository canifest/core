package dockerfile

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/benschw/opin-go/rando"
	"github.com/canifest/core/dockerfile/api"

	"github.com/canifest/core/dockerfile/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var _ = log.Print

type DockerfileTestSuite struct {
	suite.Suite

	service *DockerfileService
	host    string
}

func (suite *DockerfileTestSuite) SetupSuite() {
	host := fmt.Sprintf("localhost:%d", rando.Port())
	suite.service = &DockerfileService{}

	go suite.service.RunStandaloneClient(host)
	fmt.Println("Giving time for the service to start up.")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Ran Suite Setup.")
	suite.host = "http://" + host
}

func (suite *DockerfileTestSuite) TestAdd() {
	dockerfileClient := client.DockerfileClient{Host: suite.host}
	newDockerfile := api.Dockerfile{
		Add:  []string{"AddSting1", "AddString2"},
		Cmd:  []string{"Cmd1", "Cmd2"},
		Copy: []string{"copy1", "copy2"},
		//Dockerfile:
		Entrypoint: "Entrypoint",
		Env:        map[string]string{"key1": "value1", "key2": "value2"},
		Expose:     []string{"expose1", "epopose2"},
		From:       "From",
		//ID:
		Label: map[string]string{"labelKey1": "labelValue1",
			"labelKey2": "labelValue2"},
		Maintainer: "Maintainer: Bugs Bunney",
		Run:        []string{"Run1", "Run2"},
		User:       1004,
		Volume:     []string{"Volume1", "Volume2"},
		Workdir:    "WorkDir",
	}
	created, err := dockerfileClient.AddDockerfile(newDockerfile)
	assert.Nil(suite.T(), err)

	expected := newDockerfile
	expected.Dockerfile = "Not Yet Implemented"
	expected.ID = -1
	assert.Equal(suite.T(), expected, created)
}

func (suite *DockerfileTestSuite) TestFind() {
	dockerfileClient := client.DockerfileClient{Host: suite.host}
	expected := api.Dockerfile{
		Add:        []string{"AddSting1", "AddString2"},
		Cmd:        []string{"Cmd1", "Cmd2"},
		Copy:       []string{"copy1", "copy2"},
		Dockerfile: "Not Yet Implemented",
		Entrypoint: "Entrypoint",
		Env:        map[string]string{"key1": "value1", "key2": "value2"},
		Expose:     []string{"expose1", "epopose2"},
		From:       "From",
		ID:         -1,
		Label: map[string]string{"labelKey1": "labelValue1",
			"labelKey2": "labelValue2"},
		Maintainer: "Maintainer: Bugs Bunney",
		Run:        []string{"Run1", "Run2"},
		User:       1004,
		Volume:     []string{"Volume1", "Volume2"},
		Workdir:    "WorkDir",
	}
	found, err := dockerfileClient.FindDockerfile(-1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, found)
}

func TestDockerfileSuite(t *testing.T) {
	dockerfileSuiteTester := new(DockerfileTestSuite)
	suite.Run(t, dockerfileSuiteTester)
}
