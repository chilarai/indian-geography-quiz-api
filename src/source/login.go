package source

import (
	"config"
	"encoding/json"
	"net/http"
)

// LoginOutputData struct
type LoginOutputData struct {
	UserID    int64    `json:"userId"`
	Name    string `json:"name"`
	PhotoLink    string `json:"photoLink"`
	SessionToken string `json:"sessionToken"`
}

// LoginOutputStruct return
type LoginOutputStruct struct {
	StatusOut           config.Status  `json:"status"`
	LoginDataOut LoginOutputData `json:"data"`
}


func Login(w http.ResponseWriter, req *http.Request){
	config.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := config.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
		SessionKey string
		Name string
		PhotoLink string
	}

	var inputJSON input
	var status config.Status
	var outputStatus LoginOutputData
	var outputStruct LoginOutputStruct
	var id int = 0

	if(inputJSON.Email == "" || inputJSON.SessionKey == "" || inputJSON.Name == ""){

		status = config.Status{
			Code:    403,
			Message: config.InvalidInput,
		}
	} else {
		errUser := mydb.QueryRow("SELECT * FROM users WHERE email = ?", inputJSON.Email).Scan(&id)

		if(errUser != nil){
			status = config.Status{
				Code:    403,
				Message: errUser.Error(),
			}
		} else {

			if(id == 0){

				insUser, errIns := mydb.Exec("INSERT INTO users (email, session_key, name, photo_link) VALUES (?,?,?,?)", inputJSON.Email, inputJSON.SessionKey, inputJSON.Name, inputJSON.PhotoLink)

				if(errIns != nil){
					status = config.Status{
						Code:    403,
						Message: errIns.Error(),
					}
				} else {
					userId, _ := insUser.LastInsertId()

					outputStatus =  LoginOutputData {
						UserID: userId,
						Name: inputJSON.Name,
						PhotoLink:inputJSON.PhotoLink,
						SessionToken: inputJSON.SessionKey,
					}

					status = config.Status{
						Code:    200,
						Message: config.SuccessMsg,
					}
				}
				

			} else {

				_, errUpd := mydb.Exec("UPDATE users SET session_key = ? WHERE email = ?", inputJSON.SessionKey, inputJSON.Email)

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

	outputStruct = LoginOutputStruct{
		StatusOut: status,
		LoginDataOut: outputStatus,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
	
}