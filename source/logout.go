package source

import (
	"encoding/json"
	"indian-geography-quiz/common"
	"io/ioutil"
	"log"
	"net/http"
)

// LogoutOutputStruct return
type LogoutOutputStruct struct {
	StatusOut           common.Status  `json:"status"`
}

func Logout(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
		SessionKey string
	}

	var inputJSON input
	var status common.Status
	var outputStruct LogoutOutputStruct
	var id int = 0

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {
		if(inputJSON.Email == "" || inputJSON.SessionKey == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT id FROM users WHERE email = ? AND session_key = ?", inputJSON.Email, inputJSON.SessionKey).Scan(&id)

			if(errUser != nil){
				status = common.Status{
					Code:    403,
					Message: errUser.Error(),
				}
			} else {

				_, errUpd := mydb.Exec("UPDATE users SET session_key = ? WHERE email = ?", "", inputJSON.Email)

				if(errUpd != nil ){

					status = common.Status{
						Code:    403,
						Message: errUpd.Error(),
					}
				} else {
					status = common.Status{
						Code:    200,
						Message: common.SuccessMsg,
					}
				}					
			}
		}
	}

	outputStruct = LogoutOutputStruct{
		StatusOut: status,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}