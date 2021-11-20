package source

import (
	"common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
		Name string
		Email string
	}

	var inputJSON input
	var status common.Status
	var outputStatus LoginOutputData
	var outputStruct LoginOutputStruct
	var id int64
	var sessionKey string

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if(  inputJSON.Name == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT id, session_key, name FROM users WHERE email = ?", inputJSON.Email).Scan(&id, &sessionKey)

			newtime := time.Now().String()
			sessionHash, _ := bcrypt.GenerateFromPassword([]byte(newtime), 14)

			if(errUser != nil){

				insUser, _ := mydb.Exec("INSERT INTO users (session_key, name, email) VALUES (?,?,?)", sessionHash, inputJSON.Name, inputJSON.Email)

				userIdNew, _ := insUser.LastInsertId();

				outputStatus =  LoginOutputData {
					UserID: userIdNew,
					Name: inputJSON.Name,
					SessionToken: string(sessionHash),
				}

			} else {

				updUser, errUpd := mydb.Exec("INSERT INTO users SET name = ?, session_key = ? WHERE email = ?", inputJSON.Name, sessionHash, inputJSON.Email)

				if(errUpd != nil){
					rowsAffected, _ := updUser.RowsAffected()
					fmt.Println("Rows affected", rowsAffected)
				} else {
					outputStatus =  LoginOutputData {
						UserID: id,
						Name: inputJSON.Name,
						SessionToken: sessionKey,
					}
				}

			}

			status = common.Status{
				Code:    200,
				Message: common.SuccessMsg,
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