package main

import (
	"bytes"
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
	/*J, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}*/

	var data = &q
	var J = []byte("{\"endpoints\": [{\"path\": \"/containers/json\", \"method\": \"get\", \"summary\": \"List containers\"}, {\"path\": \"/containers/create\", \"method\": \"post\", \"summary\": \"Create a container\"}, {\"path\": \"/containers/{id}/json\", \"method\": \"get\", \"summary\": \"Inspect a container\"}, {\"path\": \"/containers/{id}/top\", \"method\": \"get\", \"summary\": \"List processes running inside a container\"}, {\"path\": \"/containers/{id}/logs\", \"method\": \"get\", \"summary\": \"Get container logs\"}, {\"path\": \"/containers/{id}/changes\", \"method\": \"get\", \"summary\": \"Get changes on a container\u2019s filesystem\"}, {\"path\": \"/containers/{id}/export\", \"method\": \"get\", \"summary\": \"Export a container\"}, {\"path\": \"/containers/{id}/stats\", \"method\": \"get\", \"summary\": \"Get container stats based on resource usage\"}, {\"path\": \"/containers/{id}/resize\", \"method\": \"post\", \"summary\": \"Resize a container TTY\"}, {\"path\": \"/containers/{id}/start\", \"method\": \"post\", \"summary\": \"Start a container\"}, {\"path\": \"/containers/{id}/stop\", \"method\": \"post\", \"summary\": \"Stop a container\"}, {\"path\": \"/containers/{id}/restart\", \"method\": \"post\", \"summary\": \"Restart a container\"}, {\"path\": \"/containers/{id}/kill\", \"method\": \"post\", \"summary\": \"Kill a container\"}, {\"path\": \"/containers/{id}/update\", \"method\": \"post\", \"summary\": \"Update a container\"}, {\"path\": \"/containers/{id}/rename\", \"method\": \"post\", \"summary\": \"Rename a container\"}, {\"path\": \"/containers/{id}/pause\", \"method\": \"post\", \"summary\": \"Pause a container\"}, {\"path\": \"/containers/{id}/unpause\", \"method\": \"post\", \"summary\": \"Unpause a container\"}, {\"path\": \"/containers/{id}/attach\", \"method\": \"post\", \"summary\": \"Attach to a container\"}, {\"path\": \"/containers/{id}/attach/ws\", \"method\": \"get\", \"summary\": \"Attach to a container via a websocket\"}, {\"path\": \"/containers/{id}/wait\", \"method\": \"post\", \"summary\": \"Wait for a container\"}, {\"path\": \"/containers/{id}/archive\", \"method\": \"get\", \"summary\": \"Get an archive of a filesystem resource in a container\"}, {\"path\": \"/containers/prune\", \"method\": \"post\", \"summary\": \"Delete stopped containers\"}, {\"path\": \"/images/json\", \"method\": \"get\", \"summary\": \"List Images\"}, {\"path\": \"/build\", \"method\": \"post\", \"summary\": \"Build an image\"}, {\"path\": \"/build/prune\", \"method\": \"post\", \"summary\": \"Delete builder cache\"}, {\"path\": \"/images/create\", \"method\": \"post\", \"summary\": \"Create an image\"}, {\"path\": \"/images/{name}/json\", \"method\": \"get\", \"summary\": \"Inspect an image\"}, {\"path\": \"/images/{name}/history\", \"method\": \"get\", \"summary\": \"Get the history of an image\"}, {\"path\": \"/images/{name}/push\", \"method\": \"post\", \"summary\": \"Push an image\"}, {\"path\": \"/images/{name}/tag\", \"method\": \"post\", \"summary\": \"Tag an image\"}, {\"path\": \"/images/search\", \"method\": \"get\", \"summary\": \"Search images\"}, {\"path\": \"/images/prune\", \"method\": \"post\", \"summary\": \"Delete unused images\"}, {\"path\": \"/auth\", \"method\": \"post\", \"summary\": \"Check auth configuration\"}, {\"path\": \"/info\", \"method\": \"get\", \"summary\": \"Get system information\"}, {\"path\": \"/version\", \"method\": \"get\", \"summary\": \"Get version\"}, {\"path\": \"/_ping\", \"method\": \"get\", \"summary\": \"Ping\"}, {\"path\": \"/commit\", \"method\": \"post\", \"summary\": \"Create a new image from a container\"}, {\"path\": \"/events\", \"method\": \"get\", \"summary\": \"Monitor events\"}, {\"path\": \"/system/df\", \"method\": \"get\", \"summary\": \"Get data usage information\"}, {\"path\": \"/images/{name}/get\", \"method\": \"get\", \"summary\": \"Export an image\"}, {\"path\": \"/images/get\", \"method\": \"get\", \"summary\": \"Export several images\"}, {\"path\": \"/images/load\", \"method\": \"post\", \"summary\": \"Import images\"}, {\"path\": \"/containers/{id}/exec\", \"method\": \"post\", \"summary\": \"Create an exec instance\"}, {\"path\": \"/exec/{id}/start\", \"method\": \"post\", \"summary\": \"Start an exec instance\"}, {\"path\": \"/exec/{id}/resize\", \"method\": \"post\", \"summary\": \"Resize an exec instance\"}, {\"path\": \"/exec/{id}/json\", \"method\": \"get\", \"summary\": \"Inspect an exec instance\"}, {\"path\": \"/volumes\", \"method\": \"get\", \"summary\": \"List volumes\"}, {\"path\": \"/volumes/create\", \"method\": \"post\", \"summary\": \"Create a volume\"}, {\"path\": \"/volumes/{name}\", \"method\": \"get\", \"summary\": \"Inspect a volume\"}, {\"path\": \"/volumes/prune\", \"method\": \"post\", \"summary\": \"Delete unused volumes\"}, {\"path\": \"/networks\", \"method\": \"get\", \"summary\": \"List networks\"}, {\"path\": \"/networks/{id}\", \"method\": \"get\", \"summary\": \"Inspect a network\"}, {\"path\": \"/networks/create\", \"method\": \"post\", \"summary\": \"Create a network\"}, {\"path\": \"/networks/{id}/connect\", \"method\": \"post\", \"summary\": \"Connect a container to a network\"}, {\"path\": \"/networks/{id}/disconnect\", \"method\": \"post\", \"summary\": \"Disconnect a container from a network\"}, {\"path\": \"/networks/prune\", \"method\": \"post\", \"summary\": \"Delete unused networks\"}, {\"path\": \"/plugins\", \"method\": \"get\", \"summary\": \"List plugins\"}, {\"path\": \"/plugins/privileges\", \"method\": \"get\", \"summary\": \"Get plugin privileges\"}, {\"path\": \"/plugins/pull\", \"method\": \"post\", \"summary\": \"Install a plugin\"}, {\"path\": \"/plugins/{name}/json\", \"method\": \"get\", \"summary\": \"Inspect a plugin\"}, {\"path\": \"/plugins/{name}/enable\", \"method\": \"post\", \"summary\": \"Enable a plugin\"}, {\"path\": \"/plugins/{name}/disable\", \"method\": \"post\", \"summary\": \"Disable a plugin\"}, {\"path\": \"/plugins/{name}/upgrade\", \"method\": \"post\", \"summary\": \"Upgrade a plugin\"}, {\"path\": \"/plugins/create\", \"method\": \"post\", \"summary\": \"Create a plugin\"}, {\"path\": \"/plugins/{name}/push\", \"method\": \"post\", \"summary\": \"Push a plugin\"}, {\"path\": \"/plugins/{name}/set\", \"method\": \"post\", \"summary\": \"Configure a plugin\"}, {\"path\": \"/nodes\", \"method\": \"get\", \"summary\": \"List nodes\"}, {\"path\": \"/nodes/{id}\", \"method\": \"get\", \"summary\": \"Inspect a node\"}, {\"path\": \"/nodes/{id}/update\", \"method\": \"post\", \"summary\": \"Update a node\"}, {\"path\": \"/swarm\", \"method\": \"get\", \"summary\": \"Inspect swarm\"}, {\"path\": \"/swarm/init\", \"method\": \"post\", \"summary\": \"Initialize a new swarm\"}, {\"path\": \"/swarm/join\", \"method\": \"post\", \"summary\": \"Join an existing swarm\"}, {\"path\": \"/swarm/leave\", \"method\": \"post\", \"summary\": \"Leave a swarm\"}, {\"path\": \"/swarm/update\", \"method\": \"post\", \"summary\": \"Update a swarm\"}, {\"path\": \"/swarm/unlockkey\", \"method\": \"get\", \"summary\": \"Get the unlock key\"}, {\"path\": \"/swarm/unlock\", \"method\": \"post\", \"summary\": \"Unlock a locked manager\"}, {\"path\": \"/services\", \"method\": \"get\", \"summary\": \"List services\"}, {\"path\": \"/services/create\", \"method\": \"post\", \"summary\": \"Create a service\"}, {\"path\": \"/services/{id}\", \"method\": \"get\", \"summary\": \"Inspect a service\"}, {\"path\": \"/services/{id}/update\", \"method\": \"post\", \"summary\": \"Update a service\"}, {\"path\": \"/services/{id}/logs\", \"method\": \"get\", \"summary\": \"Get service logs\"}, {\"path\": \"/tasks\", \"method\": \"get\", \"summary\": \"List tasks\"}, {\"path\": \"/tasks/{id}\", \"method\": \"get\", \"summary\": \"Inspect a task\"}, {\"path\": \"/tasks/{id}/logs\", \"method\": \"get\", \"summary\": \"Get task logs\"}, {\"path\": \"/secrets\", \"method\": \"get\", \"summary\": \"List secrets\"}, {\"path\": \"/secrets/create\", \"method\": \"post\", \"summary\": \"Create a secret\"}, {\"path\": \"/secrets/{id}\", \"method\": \"get\", \"summary\": \"Inspect a secret\"}, {\"path\": \"/secrets/{id}/update\", \"method\": \"post\", \"summary\": \"Update a Secret\"}, {\"path\": \"/configs\", \"method\": \"get\", \"summary\": \"List configs\"}, {\"path\": \"/configs/create\", \"method\": \"post\", \"summary\": \"Create a config\"}, {\"path\": \"/configs/{id}\", \"method\": \"get\", \"summary\": \"Inspect a config\"}, {\"path\": \"/configs/{id}/update\", \"method\": \"post\", \"summary\": \"Update a Config\"}, {\"path\": \"/distribution/{name}/json\", \"method\": \"get\", \"summary\": \"Get image information from the registry\"}, {\"path\": \"/session\", \"method\": \"post\", \"summary\": \"Initialize interactive session\"}]}")
	//Umarshalling JSON into struct
	return json.Unmarshal(J, data)
}

func AccessEndpoint(endpoint Endpoint, docker_socket string, api_version string, post_data []byte) string {
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
		response, err = httpc.Post(docker_url, "application/json", bytes.NewBuffer(post_data))
	}

	if err != nil {
		fmt.Println(err)
		response.Body.Close()
	} else {
		resp_body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(resp_body)
	}
	return ""
}

func CheckResponse(response_text string, error_msgs []string, endpoint Endpoint, test_sum string) {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	noColor := "\033[0m"

	if ContainsAny(response_text, error_msgs) && !strings.Contains(response_text, `{\"Allow\":true}`) {
		fmt.Println(endpoint.Path + " (" + endpoint.Method + ")" + test_sum + string(colorRed) + " is forbidden" + string(noColor) + " with response: " + response_text)
	} else {
		fmt.Println(endpoint.Path + " (" + endpoint.Method + ")" + test_sum + string(colorGreen) + " is allowed" + string(noColor))
	}
}

func ContainsAny(text string, list_to_check []string) bool {
	for _, check := range list_to_check {
		if strings.Contains(text, check) {
			return true
		}
	}
	return false
}

func main() {
	var error_msgs []string = []string{
		"authorization denied",
		"AuthZPlugin.AuthZReq",
		"AuthNPlugin.AuthNReq",
	}

	// Parse arguments
	help := flag.Bool("h", false, "Print help")
	error_msg := flag.String("e", "", "Indicate another error message fingerprint. Now using: "+strings.Join(error_msgs[:], ", "))
	api_version := flag.String("v", "v1.41", "Version of the docker API")

	// People shouldn't need to put a valid container id or image name and it could be dangerous
	//docker_container_id := flag.String("c", "6beb73cc1123", "Existent container ID. If not provided, false possitive regarding container actions may appear (default value is usually useless, so use this option).")
	//docker_image_name := flag.String("i", "invented_img", "Existent image name. If not provided, false possitive regarding image actions may appear (default value is usually useless, so use this option)")
	docker_container_id := "6beb73cc1123"
	docker_image_name := "invented_img32"

	flag.Parse()

	if *help || len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "[-c 6beb73cc1eef -i ubuntu [More Options]] /run/docker.sock")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(*error_msg) > 0 {
		error_msgs = append(error_msgs, *error_msg)
	}

	//Load JSON
	endpoints := &Endpoints{}
	err := endpoints.FromJSON("endpoints.json")

	if err != nil {
		panic(err)
	}

	// Check each endpoint
	for i := 0; i < len(endpoints.Endpoints); i++ {
		endpoints.Endpoints[i].Path = strings.Replace(strings.Replace(endpoints.Endpoints[i].Path, "{id}", docker_container_id, -1), "{name}", docker_image_name, -1)
		var response_text string = AccessEndpoint(endpoints.Endpoints[i], flag.Args()[0], *api_version, []byte("{}"))
		CheckResponse(response_text, error_msgs, endpoints.Endpoints[i], "")
	}

	// Check HostConfig values
	fmt.Println("\nChecking HostConfig values")
	type Post_data struct {
		Test string
		Data []byte
	}
	datas := []Post_data{
		{
			Test: " - Binds in root",
			Data: []byte(`{"Binds":{"/tmp":"/tmp"}}`),
		},
		{
			Test: " - HostConfig.Binds",
			Data: []byte(`{"HostConfig": {"Binds": ["/tmp:/tmp"]}}`),
		},
		{
			Test: " - HostConfig.Privileged",
			Data: []byte(`{"HostConfig": {"Privileged": true}}`),
		},
		{
			Test: " - HostConfig.CapAdd",
			Data: []byte(`{"HostConfig": {"CapAdd": ["SYS_ADMIN"]}}`),
		},
		{
			Test: " - HostConfig.SecurityOpt (disable apparmor)",
			Data: []byte(`{"HostConfig": {"SecurityOpt": ["apparmor:unconfined"]}}`),
		},
	}
	var endp_create Endpoint
	endp_create.Path = "/containers/create"
	endp_create.Method = "post"
	endp_create.Summary = "Create a container"

	for i := 0; i < len(datas); i++ {
		var response_text string = AccessEndpoint(endp_create, flag.Args()[0], *api_version, datas[i].Data)
		CheckResponse(response_text, error_msgs, endp_create, datas[i].Test)
	}
}
