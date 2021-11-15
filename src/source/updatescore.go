package source

import (
	"common"
	"encoding/json"
	"net/http"
	"time"
)

func UpdateScore(w http.ResponseWriter, req *http.Request){

	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		CategoryID 		int
		UserID			int
		SessionKey 		string
	}

	var inputJSON input
	var status common.Status

	if(inputJSON.SessionKey == "" || inputJSON.CategoryID == 0){

		status = common.Status{
			Code:    403,
			Message: common.InvalidInput,
		}

	} else {
		currentTime := time.Now().Format("2006-01-02")
		_, errUpdLeader := mydb.Exec("UPDATE leaderboards SET score = score + 1 WHERE user_id = ? AND score_date = ?", inputJSON.UserID, currentTime)

		if(errUpdLeader != nil){
			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			status = common.Status{
				Code:    200,
				Message: common.SuccessMsg,
			}
		}
	}

	output, _ := json.Marshal(status)

	w.Write(output)
}