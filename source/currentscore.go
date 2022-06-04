package source

import (
	"encoding/json"
	"indian-geography-quiz/common"
	"io/ioutil"
	"log"
	"net/http"
)

// CurrentScoreOutputStruct return
type CurrentScoreOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut CurrentScoreOutputData `json:"data"`
}


// CurrentScoreOutputData struct
type CurrentScoreOutputData struct {
	Score 			int
	CategoryID 		int
}

func Currentscore(w http.ResponseWriter, req *http.Request){

	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		UserID 			int
		SessionKey		string
		QuizDate 		string
	}

	var inputJSON input
	var status common.Status
	var CurrentUserScore, CurrentUserCategoryID int
	var outputStruct CurrentScoreOutputStruct
	var currentScoreStruct CurrentScoreOutputData
	
	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if(inputJSON.SessionKey == "" || inputJSON.QuizDate == "" || inputJSON.UserID == 0){
			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errSelCurrUser := mydb.QueryRow("SELECT leaderboards.score, leaderboards.quiz_id FROM leaderboards WHERE leaderboards.score_date = ? AND leaderboards.user_id =? ", inputJSON.QuizDate, inputJSON.UserID).Scan(&CurrentUserScore, &CurrentUserCategoryID)

			if(errSelCurrUser != nil){
				status = common.Status{
					Code:    403,
					Message: errSelCurrUser.Error(),
				}
			} else {
				currentScoreStruct = CurrentScoreOutputData{
					Score 			:CurrentUserScore,
					CategoryID 		:CurrentUserCategoryID,
				}

				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
			}
		}
	}

	outputStruct = CurrentScoreOutputStruct{
		StatusOut: status,
		ListDataOut: currentScoreStruct,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
	
}