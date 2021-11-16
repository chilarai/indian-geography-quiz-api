package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
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

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {
		if(inputJSON.SessionKey == "" || inputJSON.CategoryID == 0){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}

		} else {
			currentTime := time.Now().Format("2006-01-02")
			updLeader, errUpdLeader := mydb.Exec("UPDATE leaderboards SET score = score + 1 WHERE user_id = ? AND score_date = ?", inputJSON.UserID, currentTime)

			if(errUpdLeader != nil){
				status = common.Status{
					Code:    403,
					Message: errUpdLeader.Error(),
				}
			} else {
				rowsAffected, _ := updLeader.RowsAffected()

				if rowsAffected <= 0{
					_, errIns := mydb.Exec("INSERT INTO leaderboards (user_id, score, quiz_id, score_date) VALUES (?,?,?,?)", inputJSON.UserID, 1, inputJSON.CategoryID, currentTime)

					if errIns != nil {
						status = common.Status{
							Code:    403,
							Message: errIns.Error(),
						}
					} else {
						status = common.Status{
							Code:    200,
							Message: common.SuccessMsg,
						}
					}
				}  else {
					status = common.Status{
						Code:    200,
						Message: common.SuccessMsg,
					}
				}
			}
		}
	}

	output, _ := json.Marshal(status)

	w.Write(output)
}