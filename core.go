package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/canifest/core/dockerfile"
	"github.com/gorilla/mux"
)

type Status struct {
	Status string `json:"status"`
}

type Operation struct {
	Name string `json:"name"`
}

func main() {
	bindHandlers()
	printWelcomeMessageToConsole()
	startServer()
}

func bindHandlers() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	subrouter.HandleFunc("/status", statusHttpHandler).Methods("GET")
	subrouter.HandleFunc("/list", listHttpHandler).Methods("GET")
	subrouter.HandleFunc("/quit", quitHttpHandler).Methods("GET")
	subrouter.HandleFunc("/help", helpHttpHandler).Methods("GET")
	dockerfileService := dockerfile.NewDockerfileService()
	subrouter = dockerfileService.RegisterEndpoints(subrouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)
}

func operationList() []Operation {
	type Operations []Operation

	operations := Operations{
		Operation{Name: "/api/status"},
		Operation{Name: "/api/list"},
		Operation{Name: "/api/quit"},
		Operation{Name: "/api/help"},
		Operation{Name: "/api/dockerfile"},
	}

	return operations
}

func printWelcomeMessageToConsole() {
	fmt.Println("Now listening on :9993")
}

func startServer() {
	http.ListenAndServe(":9993", nil)
}

func statusHttpHandler(writer http.ResponseWriter, response *http.Request) {
	ok := Status{Status: "ok"}
	json.NewEncoder(writer).Encode(ok)
}

func listHttpHandler(writer http.ResponseWriter, response *http.Request) {
	json.NewEncoder(writer).Encode(operationList())
}

//TODO figure out how to get the writer to flush before the application shuts down
func quitHttpHandler(writer http.ResponseWriter, response *http.Request) {
	defer shutdown()
	sendHttpStatusOk(writer)
	sendByeMessageToClient(writer)
	printByeMessageToConsole()
}

func sendHttpStatusOk(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
}

func helpHttpHandler(writer http.ResponseWriter, response *http.Request) {
	sendHttpStatusOk(writer)
	sendHelpMessageToClient(writer)
}

func sendByeMessageToClient(writer http.ResponseWriter) {
	writer.Write([]byte("bye"))
}

func readHelpFile() string {
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + "/../src/core/help.txt")
	check(err)
	return string(dat)
}

func sendHelpMessageToClient(writer http.ResponseWriter) {

	var helpText string
	helpText = readHelpFile()

	json.NewEncoder(writer).Encode(helpText)
}

func printByeMessageToConsole() {
	fmt.Println("See you next time")
}

func shutdown() {
	os.Exit(0) // 0 == everything is ok
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
