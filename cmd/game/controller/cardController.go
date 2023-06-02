package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"lapson_go_api_sample/datastore/sqlc"
	"lapson_go_api_sample/metadata"
	"lapson_go_api_sample/models"
	"strconv"
	"strings"
	"time"
)

func (controller *Controller) GetCardNFT(id int) (metadata.Metadata, error) {
	result, err := controller.db.Queries.GetCardByTokenID(context.Background(), int32(id))
	var card metadata.Metadata
	if err == nil {
		card, err = controller.getCardMeta(result)
		if err != nil {
			controllerLogger.Errorln("Wrong character / equipment", id)
			controllerLogger.Errorln("err: ", err)
		}
	} else {
		controllerLogger.Errorln("Finding non exist card token id:", id)
	}
	return card, err
}

func (controller *Controller) CreateCardFromWorker(workerjson models.WorkerPostJson) (string, error) {
	name := workerjson.IpfsURI
	controllerLogger.Infoln(workerjson)
	//Get CEFile id
	ce_file_id, err := controller.db.Queries.GetCeFileByIpfsUri(context.Background(), name)
	if err != nil {
		return responseErrMsg("find ce file error", err)
	}
	ce_file_id_Record := sql.NullInt32{
		Int32: ce_file_id,
		Valid: true,
	}

	name = strings.Replace(name, ".json", "", 1)
	isUr := false
	if strings.HasPrefix(name, "UR_") {
		isUr = true
	}
	splits := strings.Split(name, "_")
	rarity := ""
	character_name := ""
	if isUr {
		rarity = splits[1]
		character_name = splits[2]
	} else if len(splits) > 2 {
		rarity = splits[0]
		character_name = splits[1] + " " + splits[2]
		if len(splits) > 3 {
			character_name = splits[1] + " " + splits[2] + " " + splits[3]
		}
	} else {
		rarity = splits[0]
		character_name = splits[1]
	}
	controllerLogger.Infoln("character name:" + character_name)

	// find rarity id
	rarity_id, err := controller.db.Queries.GetRarityIdByName(context.Background(), rarity)
	controllerLogger.Infoln("Rarity ID:")
	controllerLogger.Infoln(rarity_id)
	if err != nil {
		return responseErrMsg("find rarity error", err)
	}

	// find character / equipment id and define card parameter
	var createCard sqlc.CreateCardFromWorkerParams
	name = strings.Join(splits[:], " ")
	updateTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	var updateCard sqlc.UpdateCardFromWorkerParams
	isEquipment := false
	equipment_id, err := controller.db.Queries.GetEquipmentIDByName(context.Background(), character_name)
	if err == nil {
		isEquipment = true
	}
	if isEquipment {
		// Create Equipment
		controllerLogger.Infoln("Equipment ID:")
		controllerLogger.Infoln(equipment_id)
		equipmentIdRecord := sql.NullInt32{
			Int32: equipment_id,
			Valid: true,
		}

		createCard = sqlc.CreateCardFromWorkerParams{
			Name:         name,
			OwnerAddress: workerjson.Owner,
			EquipmentID:  equipmentIdRecord,
			Level:        1,
			TokenID:      workerjson.TokenID,
			RarityID:     rarity_id,
			IsUr:         isUr,
			CEFileID:     ce_file_id_Record,
		}

		updateCard = sqlc.UpdateCardFromWorkerParams{
			OwnerAddress: workerjson.Owner,
			TokenID:      workerjson.TokenID,
			UpdatedAt:    updateTime,
		}
	} else {
		//Create Character
		character_id, err := controller.db.Queries.GetCharacterIDByName(context.Background(), character_name)
		controllerLogger.Infoln("Character ID:")
		controllerLogger.Infoln(character_id)
		if err != nil {
			return responseErrMsg("find character / equipment error", err)
		}
		characterIdRecord := sql.NullInt32{
			Int32: character_id,
			Valid: true,
		}

		createCard = sqlc.CreateCardFromWorkerParams{
			Name:         name,
			OwnerAddress: workerjson.Owner,
			CharacterID:  characterIdRecord,
			Level:        1,
			TokenID:      workerjson.TokenID,
			RarityID:     rarity_id,
			IsUr:         isUr,
			CEFileID:     ce_file_id_Record,
		}

		updateCard = sqlc.UpdateCardFromWorkerParams{
			OwnerAddress: workerjson.Owner,
			TokenID:      workerjson.TokenID,
			UpdatedAt:    updateTime,
		}
	}

	controller.db.Mutex.RLock()
	defer controller.db.Mutex.RUnlock()

	//find exist card record or create card
	var cardRecord sqlc.Card
	existRecord, err := controller.db.Queries.GetCardByTokenID(context.Background(), int32(workerjson.TokenID))
	// if record exist update the exist record
	if err == nil {
		controllerLogger.Infoln("Updating Card")
		controllerLogger.Infoln("exist Card")
		controllerLogger.Infoln(existRecord)
		controllerLogger.Infoln("update card")
		controllerLogger.Infoln(updateCard)
		cardRecord, err = controller.db.Queries.UpdateCardFromWorker(context.Background(), updateCard)
		if err != nil {
			return responseErrMsg("update card error", err)
		}
	} else {
		controllerLogger.Infoln("Creating new Card")
		controllerLogger.Infoln(createCard)
		cardRecord, err = controller.db.Queries.CreateCardFromWorker(context.Background(), createCard)
		if err != nil {
			return responseErrMsg("create new card error", err)
		}
	}

	//created / updated card then reponse success message
	resultJson := models.ResponseSuccessJson{
		Status: "success",
		Data:   cardRecord,
	}

	resultjsonStr, err := json.MarshalIndent(resultJson, "", "\t")
	if err != nil {
		return responseErrMsg("create card: json marshal", err)
	}
	return string(resultjsonStr), nil
}

func (controller *Controller) getCardMeta(result sqlc.GetCardByTokenIDRow) (metadata.Metadata, error) {
	//replace name
	name := result.Name
	if result.CharacterID.Int32 != 0 {
		character, err := controller.db.Queries.GetCharacterByID(context.Background(), int32(result.CharacterID.Int32))
		if err == nil {
			name = character.Name
		} else {
			return metadata.Metadata{}, err
		}
	} else {
		weapon, err := controller.db.Queries.GetEquipmentByID(context.Background(), int32(result.EquipmentID.Int32))
		if err == nil {
			name = weapon.Name
		} else {
			return metadata.Metadata{}, err
		}
	}
	attribute_traits := []string{
		//"UR",
		"Rarity",
		"Level",
		"Type",
	}
	var attributes []metadata.Attribute
	filename := result.ImageUri
	for _, trait := range attribute_traits {
		var value string
		if trait == "Level" {
			value = strconv.Itoa(int(result.Level))
		}
		if trait == "Rarity" {
			if result.Rarity.Valid {
				value = result.Rarity.String
			}
		}
		if trait == "Type" {
			if result.CharacterID.Int32 != 0 {
				value = "Character"
			} else {
				value = "Weapon"
			}
		}

		attributes = append(attributes,
			metadata.Attribute{
				TraitType: trait,
				Value:     value,
			},
		)
	}
	filename_image := controller.allconfig.MetaJasonConfig.IpfsImgURL
	filename_image = filename_image + filename
	// if strings.HasPrefix(filename, "UR_") {
	// 	filename_image = filename_image + filename + ".mp4"
	// } else {
	// 	filename_image = filename_image + filename + ".jpg"
	// }
	//filename_json := filename + ".json"
	exturnal_url := controller.allconfig.MetaJasonConfig.ExternalURL

	card := metadata.Metadata{
		Name:        name,
		Image:       filename_image,
		ExternalURL: exturnal_url,
		Attributes:  attributes,
	}
	return card, nil
}
