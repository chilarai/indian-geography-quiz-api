package source

import (
	"common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	StatusOut           common.Status  `json:"status"`
	LoginDataOut LoginOutputData `json:"data"`
}


func Login(w http.ResponseWriter, req *http.Request){
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
		SessionKey string
		Name string
		PhotoLink string
	}

	var inputJSON input
	var status common.Status
	var outputStatus LoginOutputData
	var outputStruct LoginOutputStruct
	var id int64

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if(inputJSON.Email == "" || inputJSON.SessionKey == "" || inputJSON.Name == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT id FROM users WHERE email = ?", inputJSON.Email).Scan(&id)

			if(errUser != nil){
				fmt.Println(errUser.Error())
				insUser, errIns := mydb.Exec("INSERT INTO users (email, session_key, name, photo_link) VALUES (?,?,?,?)", inputJSON.Email, inputJSON.SessionKey, inputJSON.Name, inputJSON.PhotoLink)

				if(errIns != nil){
					status = common.Status{
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

					status = common.Status{
						Code:    200,
						Message: common.SuccessMsg,
					}
				}
			} else {

				_, errUpd := mydb.Exec("UPDATE users SET session_key = ? WHERE email = ?", inputJSON.SessionKey, inputJSON.Email)

				if(errUpd != nil ){

					status = common.Status{
						Code:    403,
						Message: errUpd.Error(),
					}
				} else {
					outputStatus =  LoginOutputData {
						UserID: id,
						Name: inputJSON.Name,
						PhotoLink:inputJSON.PhotoLink,
						SessionToken: inputJSON.SessionKey,
					}

					status = common.Status{
						Code:    200,
						Message: common.SuccessMsg,
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