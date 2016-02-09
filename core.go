package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	http.HandleFunc("/status", statusHttpHandler)
	http.HandleFunc("/list", listHttpHandler)
	http.HandleFunc("/quit", quitHttpHandler)
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

func operationList() []Operation {
	type Operations []Operation

	operations := Operations{
		Operation{Name: "/status"},
		Operation{Name: "/list"},
		Operation{Name: "/quit"},
	}

	return operations
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

func sendByeMessageToClient(writer http.ResponseWriter) {
	writer.Write([]byte("bye"))
}

func printByeMessageToConsole() {
	fmt.Println("See you next time")
}

func shutdown() {
	os.Exit(0) // 0 == everything is ok
}
