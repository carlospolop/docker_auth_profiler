package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Summary string `json:"summary"`
}

func (q *Endpoints) FromJSON(file string) error {
	//Reading JSON file
	J, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var data = &q
	//Umarshalling JSON into struct
	return json.Unmarshal(J, data)
}

func access_endpoint(endpoint Endpoint, docker_socket string, api_version string) string {
	var err error
	httpc := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", docker_socket)
			},
		},
	}
	var response *http.Response
	var docker_url = "http://localhost/" + api_version + endpoint.Path
	if endpoint.Method == "get" {
		response, err = httpc.Get(docker_url)
	} else {
		response, err = httpc.Post(docker_url, "application/json", strings.NewReader("{}"))
	}

	if err != nil {
		panic(err)
	}

	resp_body, _ := ioutil.ReadAll(response.Body)
	return string(resp_body)
}

func check_response(response_text string, error_msg string, endpoint Endpoint, print_all bool) {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	noColor := "\033[0m"

	if strings.Contains(response_text, error_msg) && print_all {
		fmt.Println(endpoint.Path + "(" + endpoint.Method + ")" + string(colorRed) + " is forbidden" + string(noColor))
	} else {
		fmt.Println(endpoint.Path + "(" + endpoint.Method + ")" + string(colorGreen) + " is allowed" + string(noColor))
	}
}

func main() {
	// Parse arguments
	help := flag.Bool("h", false, "Print help")
	print_all := flag.Bool("a", false, "Print allowed and forbidden")
	error_msg := flag.String("e", "authorization denied by plugin", "Indicate the error message fingerprint")
	api_version := flag.String("v", "v1.41", "Version of the docker API")
	flag.Parse()

	if *help || len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "/run/docker.sock")
		flag.PrintDefaults()
		os.Exit(0)
	}

	//Load JSON
	endpoints := &Endpoints{}
	err := endpoints.FromJSON("endpoints.json")

	if err != nil {
		panic(err)
	}

	// Check each endpoint
	for i := 0; i < len(endpoints.Endpoints); i++ {
		var response_text string = access_endpoint(endpoints.Endpoints[i], flag.Args()[0], *api_version)
		check_response(response_text, *error_msg, endpoints.Endpoints[i], *print_all)
	}
}
