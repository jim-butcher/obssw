package util

import (
	"errors"
	"fmt"
	"log"

	"obsws"
)

func GetCurrentSceneName(client obsws.Client) string {
	resp, err := obsws.NewGetCurrentSceneRequest().SendReceive(client)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Name
}

func GetCameraSourceName(client obsws.Client) (string, error) {
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
