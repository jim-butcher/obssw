package main

import (
	"log"
	"os"

	"internal/util"

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

	var cameraName string
	if len(os.Args) == 2 {
		cameraName = os.Args[1]
	} else {
		cameraName, err = util.GetCameraSourceName(client)
	}

	sceneName := util.GetCurrentSceneName(client)
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
