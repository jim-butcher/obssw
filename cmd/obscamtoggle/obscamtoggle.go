package main

import (
	"errors"
	"fmt"
	"log"

	obsws "github.com/christopher-dG/go-obs-websocket"
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

	sceneName := getCurrentSceneName(client)
	cameraName, err := getCameraSourceName(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := obsws.NewGetSceneItemPropertiesRequest(sceneName, cameraName).SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}

	_, err = obsws.NewSetSceneItemPropertiesRequest(
		sceneName,
		cameraName,
		resp.PositionX,
		resp.PositionY,
		resp.PositionAlignment,
		resp.Rotation,
		resp.ScaleX,
		resp.ScaleY,
		resp.CropTop,
		resp.CropBottom,
		resp.CropLeft,
		resp.CropRight,
		!resp.Visible,
		resp.Locked,
		resp.BoundsType,
		resp.BoundsAlignment,
		resp.BoundsX,
		resp.BoundsY,
	).SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentSceneName(client obsws.Client) string {
	resp, err := obsws.NewGetCurrentSceneRequest().SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Name
}

func getCameraSourceName(client obsws.Client) (string, error) {
	resp, err := obsws.NewGetSourcesListRequest().SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}

	for _, source := range resp.Sources {
		if source["typeId"] == "v4l2_input" {
			return fmt.Sprintf("%v", source["name"]), nil
		}
	}

	return "", errors.New("could not find camera source")
}
