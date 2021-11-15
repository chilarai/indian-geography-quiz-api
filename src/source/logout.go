package source

import (
	"config"
	"encoding/json"
	"net/http"
)

// LogoutOutputStruct return
type LogoutOutputStruct struct {
	StatusOut           config.Status  `json:"status"`
}

func Logout(w http.ResponseWriter, req *http.Request){
	
	config.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := config.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
		SessionKey string
	}

	var inputJSON input
	var status config.Status
	var outputStruct LogoutOutputStruct
	var id int = 0

	if(inputJSON.Email == "" || inputJSON.SessionKey == ""){

		status = config.Status{
			Code:    403,
			Message: config.InvalidInput,
		}
	} else {
		errUser := mydb.QueryRow("SELECT * FROM users WHERE email = ? AND session_key = ?", inputJSON.Email, inputJSON.SessionKey).Scan(&id)

		if(errUser != nil){
			status = config.Status{
				Code:    403,
				Message: errUser.Error(),
			}
		} else {

			if(id != 0){

				_, errUpd := mydb.Exec("UPDATE users SET session_key = ? WHERE email = ?", "", inputJSON.Email)

				if(errUpd != nil ){

					status = config.Status{
						Code:    403,
						Message: errUpd.Error(),
					}
				} else {
					status = config.Status{
						Code:    200,
						Message: config.SuccessMsg,
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