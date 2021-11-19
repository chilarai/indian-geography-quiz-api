package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// LoginOutputData struct
type LoginOutputData struct {
	UserID    int64    `json:"userId"`
	Name    string `json:"name"`
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
		Password string
		Email string
	}

	var inputJSON input
	var status common.Status
	var outputStatus LoginOutputData
	var outputStruct LoginOutputStruct
	var id int64
	var password, sessionKey, name string

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if( inputJSON.Password == "" || inputJSON.Email == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT id, password, name, session_key FROM users WHERE email = ?", inputJSON.Email).Scan(&id, &password, &name, &sessionKey)

			if(errUser != nil){
				status = common.Status{
					Code:    403,
					Message: errUser.Error(),
				}
			} else {

				passHash, _ := bcrypt.GenerateFromPassword([]byte(inputJSON.Password), 14)
				if(password == string(passHash)){
					
					status = common.Status{
						Code:    403,
						Message: common.FailedMsg,
					}

				} else {

					outputStatus =  LoginOutputData {
						UserID: id,
						Name: name,
						SessionToken: sessionKey,
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