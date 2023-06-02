package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"lapson_go_api_sample/datastore/sqlc"
	"lapson_go_api_sample/models"
	"time"
)

func (controller *Controller) SaveMiniGameMark(wallet_address string, mark int32) (string, error) {
	exist_record, err := controller.db.Queries.GetMiniGameRankByWallet(context.Background(), wallet_address)
	updateTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	var miniGameMark sqlc.MinigameRank
	controller.db.Mutex.RLock()
	defer controller.db.Mutex.RUnlock()
	if err == nil {
		if mark > exist_record.Mark {
			//update current record
			updateMark := sqlc.UpdateMiniGameRankParams{
				Mark:          mark,
				UpdatedAt:     updateTime,
				WalletAddress: wallet_address,
			}
			controllerLogger.Infoln("Updating Existing MiniGame mark record")
			controllerLogger.Infoln(exist_record)
			controllerLogger.Infoln("updated mark")
			controllerLogger.Infoln(updateMark)
			miniGameMark, err = controller.db.Queries.UpdateMiniGameRank(context.Background(), updateMark)
			if err != nil {
				return responseErrMsg("update miniGame  error", err)
			}
		} else {
			err := errors.New("Current score lower than existing record")
			return responseErrMsg("update miniGame  error: ", err)
		}
	} else {
		//create new mini game record
		createMark := sqlc.CreateMiniGameRankParams{
			WalletAddress: wallet_address,
			Mark:          mark,
		}
		controllerLogger.Infoln("Creating MiniGame mark record")
		controllerLogger.Infoln(createMark)
		miniGameMark, err = controller.db.Queries.CreateMiniGameRank(context.Background(), createMark)
		if err != nil {
			return responseErrMsg("create miniGame mark error", err)
		}
	}

	resultJson := models.ResponseSuccessJson{
		Status: "success",
		Data:   miniGameMark,
	}

	resultjsonStr, err := json.MarshalIndent(resultJson, "", "\t")
	if err != nil {
		return responseErrMsg("create miniGame Mark: json marshal", err)
	}
	return string(resultjsonStr), nil
}

func (controller *Controller) GetMiniGameMarks() (string, error) {
	mark_records, err := controller.db.Queries.GetMiniGameRank(context.Background())
	if err == nil {
		mark_record_result := []models.MiniGamePostJson{}
		for _, mark_record := range mark_records {
			mini_game_json := models.MiniGamePostJson{
				Wallet: mark_record.WalletAddress,
				Mark:   mark_record.Mark,
			}
			mark_record_result = append(mark_record_result, mini_game_json)
		}
		resultJson := models.ResponseSuccessJson{
			Status: "success",
			Data:   mark_record_result,
		}
		resultjsonStr, err := json.MarshalIndent(resultJson, "", "\t")
		if err != nil {
			return responseErrMsg("Get miniGame Mark: json marshal", err)
		}
		return string(resultjsonStr), nil
	} else {
		return responseErrMsg("Get miniGame mark error", err)
	}
}

type RankData struct {
	Rank int64 `json:"rank"`
	Mark int32 `json:"mark"`
}

func (controller *Controller) GetMiniGameRankByWallet(wallet_address string) (string, error) {
	rank, err := controller.db.Queries.GetRankByWallet(context.Background(), wallet_address)
	controllerLogger.Infoln(rank)
	if err == nil {
		rank_data := RankData{
			Rank: rank.Rank,
			Mark: rank.Mark,
		}
		resultJson := models.ResponseSuccessJson{
			Status: "success",
			Data:   rank_data,
		}
		resultjsonStr, err := json.MarshalIndent(resultJson, "", "\t")
		if err != nil {
			return responseErrMsg("Get miniGame Rank: json marshal", err)
		}
		return string(resultjsonStr), nil
	} else {
		return responseErrMsg("Get Rank error", err)
	}
}

func (controller *Controller) GetMiniGameTop10() (string, error) {
	top10scores, err := controller.db.Queries.GetTop10Scores(context.Background())
	if err == nil {
		resultJson := models.ResponseSuccessJson{
			Status: "success",
			Data:   top10scores,
		}
		resultjsonStr, err := json.MarshalIndent(resultJson, "", "\t")
		if err != nil {
			return responseErrMsg("Get miniGame Top10 failed: json marshal", err)
		}
		return string(resultjsonStr), nil
	} else {
		return responseErrMsg("Get miniGame Top10 failed error: ", err)
	}
}
