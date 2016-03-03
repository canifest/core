package dockerfile

import (
	"fmt"
	"net/http"

	"github.com/benschw/opin-go/rest"
	"github.com/canifest/core/dockerfile/api"
)

type DockerfileResource struct{}

func (r *DockerfileResource) Add(res http.ResponseWriter, req *http.Request) {
	var dockerfile api.Dockerfile
	if err := rest.Bind(req, &dockerfile); err != nil {
		rest.SetBadRequestResponse(res)
	}

	//generate and store dockerfile here.
	dockerfile.Dockerfile = "Not Yet Implemented"
	dockerfile.ID = -1

	rest.SetCreatedResponse(res, dockerfile, fmt.Sprintf("dockerfile/%d", dockerfile.ID))
}

func (r *DockerfileResource) Find(res http.ResponseWriter, req *http.Request) {
	foundDockerfile := api.Dockerfile{
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

	rest.SetOKResponse(res, foundDockerfile)
}
