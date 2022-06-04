package source

import (
	"encoding/json"
	"indian-geography-quiz/common"
	"io/ioutil"
	"log"
	"net/http"
)

// LeaderboardsOutputStruct return
type LeaderboardsOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut []*LeaderboardsOutputData `json:"data"`
}

// LeaderboardsOutputData struct
type LeaderboardsOutputData struct {
	UserID			int
	Score 			int
	CategoryID 		int
	Name            string
}

func Leaderboard(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		UserID 			int
		CategoryID 		int
		SessionKey 		string
		QuizDate 		string
	}

	var inputJSON input
	var Name string
	var UserID, Score, CategoryID int

	var status common.Status
	var outputStruct LeaderboardsOutputStruct
	var result []*LeaderboardsOutputData

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if(inputJSON.SessionKey == "" || inputJSON.QuizDate == ""){
			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			selLeaderboards, errSel := mydb.Query("SELECT users.id, users.name, leaderboards.score, leaderboards.quiz_id FROM leaderboards LEFT JOIN users ON users.id = leaderboards.user_id WHERE leaderboards.score_date = ? ORDER BY leaderboards.score DESC LIMIT 0, 20", inputJSON.QuizDate)

			if(errSel != nil){

				status = common.Status{
					Code:    403,
					Message: errSel.Error(),
				}

			} else{
				for selLeaderboards.Next(){
					errSelScan := selLeaderboards.Scan(&UserID, &Name, &Score, &CategoryID)

					if errSelScan != nil {
						log.Println(errSelScan)
					} else {

						elem := LeaderboardsOutputData{
							UserID			:UserID,
							Score 			:Score,
							CategoryID 		:CategoryID,
							Name            :Name,
						}

						result = append(result, &elem)
					}
				}

				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
			}
		}
	}

	outputStruct = LeaderboardsOutputStruct{
		StatusOut: status,
		ListDataOut: result,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}