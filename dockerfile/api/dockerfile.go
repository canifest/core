package api

type Dockerfile struct {
	Add        []string          `json:"add"`
	Cmd        []string          `json:"cmd"`
	Copy       []string          `json:"copy"`
	Entrypoint string            `json:"entrypoint"`
	Dockerfile string            `json:"dockerfile"`
	Env        map[string]string `json:"env"`
	Expose     []string          `json:"expose"`
	From       string            `json:"from"`
	ID         int               `json:"id"`
	Label      map[string]string `json:"label"`
	Maintainer string            `json:"maintainer"`
	Run        []string          `json:"run"`
	User       int               `json:"user"`
	Volume     []string          `json:"volume"`
	Workdir    string            `json:"workdir"`
}
