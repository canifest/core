package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

type Dockerfile struct {
    Add []string `json:"add"`
    Cmd []string `json:"cmd"`
    Copy []string `json:"copy"`
    Entrypoint string `json:"entrypoint"`
    Env map[string]string `json:"env"`
    Expose []string `json:"expose"`
    From string `json:"from"`
    Label map[string]string `json:"label"`
    Maintainer string `json:"maintainer"`
    Run []string `json:"run"`
    User int `json:"user"`
    Volume []string `json:"volume"`
    Workdir string `json:"workdir"`
}

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
	http.HandleFunc("/status", statusHttpHandler)
	http.HandleFunc("/list", listHttpHandler)
	http.HandleFunc("/quit", quitHttpHandler)
	http.HandleFunc("/help", helpHttpHandler)
	http.HandleFunc("/dockerfile", dockerfileHttpHandler)
}

func operationList() []Operation {
	type Operations []Operation

	operations := Operations{
		Operation{Name: "/status"},
		Operation{Name: "/list"},
		Operation{Name: "/quit"},
		Operation{Name: "/help"},
		Operation{Name: "/dockerfile"},
	}

	return operations
}

func generateDockerfile() Dockerfile {

	dockerfile := Dockerfile{Entrypoint: "/here"}

	return dockerfile
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

func dockerfileHttpHandler(writer http.ResponseWriter, response *http.Request) {
	json.NewEncoder(writer).Encode(generateDockerfile())
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
	dat, err := ioutil.ReadFile(pwd+"/../src/core/help.txt")
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
