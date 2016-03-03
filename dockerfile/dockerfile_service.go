package dockerfile

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DockerfileService struct {
}

func NewDockerfileService() *DockerfileService {
	s := &DockerfileService{}
	return s
}

func (s *DockerfileService) RunStandaloneClient(bind string) error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	subrouter = s.RegisterEndpoints(subrouter)
	http.Handle("/", router)
	fmt.Println("Running Standalone Dockerclient at " + bind)
	return http.ListenAndServe(bind, nil)
}

func (s *DockerfileService) RegisterEndpoints(router *mux.Router) *mux.Router {
	//route handlers
	resource := &DockerfileResource{}
	router.HandleFunc("/dockerfile", resource.Add).Methods("POST")
	router.HandleFunc("/dockerfile/{id}", resource.Find).Methods("GET")
	//router.HandleFunc("/dockerfile/{id}", resource.Update).Methods("PUT")
	//router.HandleFunc("/dockerfile/{id}", resource.Delete).Methods("DELETE")
	return router
}
