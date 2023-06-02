package controller

import (
	"context"
	"lapson_go_api_sample/datastore/sqlc"
	"lapson_go_api_sample/metadata"
	"strconv"
	"strings"
)

func (controller *Controller) GetBuilding(id int) (metadata.Metadata, error) {
	result, err := controller.db.Queries.GetBuilding(context.Background(), int32(id))
	var building metadata.Metadata
	if err == nil {
		building, _ = controller.getBuildingMeta(result)
	}
	return building, err
}

func (controller *Controller) GetListBuilding() ([]metadata.Metadata, error) {
	results, err := controller.db.Queries.ListBuilding(context.Background())
	var buildings []metadata.Metadata

	for _, result := range results {
		building, _ := controller.getBuildingMeta(result)
		buildings = append(buildings, building)
		//generateJsonFile(building, filename_json)

	}
	return buildings, err
}

func (controller *Controller) getBuildingMeta(result sqlc.Building) (metadata.Metadata, string) {
	attribute_traits := []string{
		"UR",
		"Rarity",
		"Level",
	}
	var attributes []metadata.Attribute
	filename := getFileName(result.Name)
	for _, trait := range attribute_traits {
		var trait_value string
		trait_value = getCardField(&result, trait)
		if trait == "UR" {
			if result.Isur {
				trait_value = "Yes"
			} else {
				trait_value = "No"
			}
		}
		if trait == "Level" {
			trait_value = strconv.Itoa(int(result.Level))
		}

		attributes = append(attributes,
			metadata.Attribute{
				TraitType: trait,
				Value:     trait_value,
			},
		)
	}
	filename_image := controller.allconfig.MetaJasonConfig.BuildingIpfsImageURL
	if strings.HasPrefix(filename, "UR_") {
		filename_image = filename_image + filename + ".mp4"
	} else {
		filename_image = filename_image + filename + ".jpg"
	}
	filename_json := filename + ".json"
	exturnal_url := controller.allconfig.MetaJasonConfig.ExternalURL

	building := metadata.Metadata{
		Name:        result.Name,
		Image:       filename_image,
		ExternalURL: exturnal_url,
		Attributes:  attributes,
	}
	return building, filename_json
}
