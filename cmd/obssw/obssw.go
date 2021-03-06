package main

import (
	"log"
	"os"
	"strconv"

	"obsws"
)

func main() {
	// connect to OBS
	client := obsws.Client{
		Host:     "localhost",
		Port:     4444,
		Password: "foobar",
	}
	err := client.Connect()
	if err != nil {
		log.Fatal("Could not connect to obs websocket")
	}
	defer client.Disconnect()
	log.Println("Connected to OBS websocket")

	sceneIdx, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// get scene name from the list of scenes
	resp, err := obsws.NewGetSceneListRequest().SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}

	sceneName := resp.Scenes[sceneIdx].Name

	// set current scene by index
	_, err = obsws.NewSetCurrentSceneRequest(sceneName).SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}
}
